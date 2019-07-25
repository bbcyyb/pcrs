package database

import (
	"database/sql"
	"fmt"
	"github.com/bbcyyb/pcrs/conf"
	"time"

	"github.com/bbcyyb/pcrs/pkg/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type Connection struct {
	Db *gorm.DB
}

const (
	dbType = "mssql"
)

var DbConnection *Connection

// SQLCommon is the minimal database connection functionality.
type SQLCommon interface {
	QuerySingle(result interface{}, query string, args ...interface{}) bool
	QueryMany(resList []interface{}, query string, args ...interface{}) bool
	Exec(query string, args ...interface{}) (sql.Result, bool)
	ExecTx(query string, args ...interface{}) (sql.Result, bool)
}

func Setup() {
	DbConnection = NewConnection()
}

func NewConnection() *Connection {
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
	logger.Log.Info("connect mssql ....")

	return &Connection{db}
}

func (c *Connection) CloseDB() {
	defer c.Db.Close()
}

func (c *Connection) query(query string, args ...interface{}) (*sql.Rows, error) {
	return c.Db.DB().Query(query, args)
}

func (c *Connection) queryRow(query string, args ...interface{}) *sql.Row {
	return c.Db.DB().QueryRow(query, args)
}

func (c *Connection) exec(query string, args ...interface{}) (sql.Result, error) {
	return c.Db.DB().Exec(query, args)
}

// Query single
func (c *Connection) QuerySingle(result interface{}, query string, args ...interface{}) bool {
	err := c.queryRow(query, args).Scan(&result)
	if err != nil {
		logger.Log.Errorf("occur query error, with sql:%v, error:%v", query, err)
		return false
	}
	return true
}

// Query Many
func (c *Connection) QueryMany(resList []interface{}, query string, args ...interface{}) bool {
	rows, err := c.query(query, args)
	if err != nil {
		logger.Log.Errorf("occur query error, with sql:%v, error:%v", query, err)
		return false
	}
	defer rows.Close()
	for rows.Next() {
		var res interface{}
		err = c.Db.ScanRows(rows, &res)
		if err != nil {
			logger.Log.Errorf("scanRows error, with sql:%v, error:%v", query, err)
			return false
		}
		resList = append(resList, res)
	}
	return true
}

// Create/Update/Delete
func (c *Connection) Exec(query string, args ...interface{}) (sql.Result, bool) {
	res, err := c.exec(query, args)
	if err != nil {
		logger.Log.Errorf("fail to execute sql:%v， error:%v", query, err)
		return nil, false
	}
	return res, true
}

func (c *Connection) ExecTx(query string, args ...interface{}) (sql.Result, bool) {
	tx := c.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		logger.Log.Errorf("create tx failed, with error:%v", err)
		return nil, false
	}

	res, err := tx.DB().Exec(query, args)
	if err != nil {
		tx.Rollback()
		logger.Log.Errorf("execute tx failed, with error:%v", err)
		return nil, false
	}
	err = tx.Commit().Error
	if err != nil {
		logger.Log.Errorf("commit tx failed, with error:%v", err)
		return nil, false
	}
	return res, true
}
