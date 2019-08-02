package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/bbcyyb/pcrs/conf"
	"reflect"
	"time"

	"github.com/bbcyyb/pcrs/pkg/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

// SQLCommon is the minimal database connection functionality.
type SQLCommon interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	// Transactional execution
	ExecTx(query string, args ...interface{}) (sql.Result, error)
	QuerySingle(result interface{}, query string, args ...interface{}) error
	QueryMany(resList []interface{}, query string, args ...interface{}) error
}

// SqlRows is the sql.Rows interface
type SqlRows interface {
	Next() bool
	Scan(...interface{}) error
	Columns() ([]string, error)
	Close() error
}

type Conn struct {
	Db   *gorm.DB
	rows SqlRows
}

const (
	dbType = "mssql"
)

var DbConn *Conn

func Setup() {
	DbConn = NewConnection()
}

func NewConnection() *Conn {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		conf.C.DB.Server, conf.C.DB.User, conf.C.DB.Password, conf.C.DB.Port, conf.C.DB.Database)
	db, err := gorm.Open(dbType, connectionString)
	if err != nil {
		logger.Log.Errorf("connect string:%v, dbType:%v,  err:%v", connectionString, dbType, err)
		panic("failed to connect database")
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
	err = db.DB().Ping()
	if err != nil {
		panic("failed to connect database")
	}
	logger.Log.Info("connect mssql successfully")

	return &Conn{db, nil}
}

func (c *Conn) Close() {
	if c.rows != nil {
		defer c.rows.Close()
	}
	defer c.Db.Close()
}

func (c *Conn) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return c.Db.DB().Query(query, args...)
}

func (c *Conn) QueryRow(query string, args ...interface{}) *sql.Row {
	return c.Db.DB().QueryRow(query, args...)
}

// Query single
func (c *Conn) QuerySingle(result interface{}, query string, args ...interface{}) error {
	rows, err := c.Query(query, args...)
	if err != nil {
		logger.Log.Errorf("occur Query error, with sql:%v, error:%v", query, err)
		return err
	}
	defer rows.Close()
	c.rows = rows
	c.scanRow(result)
	if err != nil {
		logger.Log.Errorf("QuerySingle error, with sql:%v, error:%v", query, err)
		return err
	}
	return nil
}

// Query Many
func (c *Conn) QueryMany(resList interface{}, query string, args ...interface{}) error {
	rows, err := c.Query(query, args...)
	if err != nil {
		logger.Log.Errorf("occur Query error, with sql:%v, error:%v", query, err)
		return err
	}
	defer rows.Close()
	c.rows = rows
	c.scanRow(resList)
	if err != nil {
		logger.Log.Errorf("QueryMany error, with sql:%v, error:%v", query, err)
		resList = nil
		return err
	}
	return nil
}

// Create/Update/Delete
func (c *Conn) Exec(query string, args ...interface{}) (sql.Result, error) {
	res, err := c.exec(query, args...)
	if err != nil {
		logger.Log.Errorf("fail to execute sql:%v， error:%v", query, err)
		return nil, err
	}
	return res, nil
}

func (c *Conn) ExecTx(query string, args ...interface{}) (sql.Result, error) {
	tx := c.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		logger.Log.Errorf("create tx failed, with error:%v", err)
		return nil, err
	}

	res, err := tx.CommonDB().Exec(query, args...)
	if err != nil {
		tx.Rollback()
		logger.Log.Errorf("execute tx failed, with error:%v", err)
		return nil, err
	}
	err = tx.Commit().Error
	if err != nil {
		logger.Log.Errorf("commit tx failed, with error:%v", err)
		return nil, err
	}
	return res, nil
}

func (c *Conn) exec(query string, args ...interface{}) (sql.Result, error) {
	return c.Db.DB().Exec(query, args...)
}

func (c *Conn) scan(target interface{}, single bool) error {
	targetValue := reflect.ValueOf(target)

	if targetValue.Kind() != reflect.Ptr {
		return errors.New("not a pointer")
	}

	targetValue = targetValue.Elem()

	if !single && targetValue.Kind() != reflect.Slice {
		return errors.New("not a slice")
	}

	if single && targetValue.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}

	valueType := targetValue.Type()

	results := make(map[string][]interface{})

	cols, err := c.rows.Columns()
	if err != nil {
		return err
	}

	rows := 0
	for c.rows.Next() {
		rows++
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := c.rows.Scan(columnPointers...); err != nil {
			return err
		}

		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			results[colName] = append(results[colName], *val)
		}
	}

	for i := 0; i < rows; i++ {
		var elemType reflect.Type
		if single {
			elemType = valueType
		} else {
			elemType = valueType.Elem()
		}
		num := elemType.NumField()
		elemNew := reflect.New(elemType).Elem()

		for j := 0; j < num; j++ {
			field := elemType.Field(j).Tag.Get("json")
			if field == "" {
				continue
			}
			value := results[field][i]
			switch value.(type) {
			case int64:
				elemNew.Field(j).SetInt(value.(int64))
			default:
				tmp := reflect.ValueOf(value)
				elemNew.Field(j).Set(tmp)
			}
		}
		if single {
			targetValue.Set(elemNew)
			break
		} else {
			targetValue.Set(reflect.Append(targetValue, elemNew))
		}
	}
	return nil
}

func (c *Conn) scanRows(target interface{}) error {
	err := c.scan(target, false)
	if err != nil {
		return err
	}
	return nil
}

func (c *Conn) scanRow(target interface{}) error {
	err := c.scan(target, true)
	if err != nil {
		return err
	}
	return nil
}
