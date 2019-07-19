package database

import (
	"fmt"
	"time"

	"github.com/bbcyyb/pcrs/pkg/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type Connection struct {
	Db *gorm.DB
}

var DbConnection *Connection

func init() {
	DbConnection = Setup()
}

func Setup() *Connection {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
		port                                              int
	)
	dbType = "mssql"
	dbName = "PowerCalc"
	user = "PowerCalc"
	password = "Power@1433"
	port = 1433
	host = "10.35.83.61"
	tablePrefix = "blog_"

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		host, user, password, port, dbName)
	db, err := gorm.Open(dbType, connectionString)

	if err != nil {
		log.Error(err)
		panic("failed to connect database")
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	log.Info("connect mssql ....")

	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// db.Callback().Create().Get("gorm:create")
	// db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// db.Callback().Query()
	return &Connection{db}
}

func CloseDB() {
	defer DbConnection.Db.Close()
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifyTime` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}
