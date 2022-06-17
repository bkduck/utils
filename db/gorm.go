package db

import (
	"fmt"
	"gorm.io/gorm/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
	Tx   *gorm.DB
	Data interface{}
}

func (sqlDb *Db) Begins() {
	sqlDb.Tx = sqlDb.DB.Begin()
}

func (sqlDb *Db) Commits() error {
	err := sqlDb.Tx.Commit().Error
	sqlDb.Tx = nil
	return err
}

func (sqlDb *Db) Rollbacks() error {
	err := sqlDb.Tx.Rollback().Error
	sqlDb.Tx = nil
	return err
}

type dialOptions func(db *gorm.DB)

/*
func Dial(host string, options ...dialOptions) *gorm.DB {
	con, err := gorm.Open("mysql", host)
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, arg = %v the error is '%v'", host, err))
	}
	//设置表名前缀
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return defaultTableName[:]
	//}
	//con.DB().SetConnMaxLifetime(2 * time.Hour)
	for _, option := range options {
		option(con)
	}
	return con
}


func DialAutoMigrate(value ...interface{}) dialOptions {
	return func(db *gorm.DB) {
		db.AutoMigrate(value...)
	}
}
*/
/*
func DialLogMode(enable bool) dialOptions {
	return func(db *gorm.DB) {
		db.LogMode(enable)
	}
}

func DialMaxCon(maxCon int) dialOptions {
	return func(db *gorm.DB) {
		db.DB().SetMaxOpenConns(maxCon)
		idle := maxCon
		if maxCon/3 >= 10 {
			idle = maxCon / 3
		}
		db.DB().SetMaxIdleConns(idle)
	}
}
*/


func NewMysql(args string, maxCon int, arr []interface{}, enable bool) *gorm.DB {
	var con *gorm.DB
	var err error
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	con, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger:logger.Default.LogMode(logger.Info), //开启sql日志
	})
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, arg = %v the error is '%v'", args, err))
	}
	//设置表名前缀
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return defaultTableName[:]
	//}
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return defaultTableName[:len(defaultTableName)-1]
	//}

	//开启sql日志
	con.Logger.LogMode(logger.Info)

	sqlDB, err := con.DB()
	if err != nil {
		panic(fmt.Sprintf("Got error when get con.DB, arg = %v the error is '%v'", args, err))
	}

	idle := maxCon
	if maxCon/3 >= 10 {
		idle = maxCon / 3
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(idle)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(2 * time.Hour)

	//若结构有变，则删除表重新创建
	//dropTable(con, arr...)
	//con.AutoMigrate(arr...) //若没有表，自动生成表
	return con
}

func (sqlDb *Db) Create() (error, interface{}) {
	//var err error
	result := sqlDb.DB.Create(sqlDb.Data)
	return result.Error, result.RowsAffected
}

func Create(db *gorm.DB, value interface{}) (error, interface{}) {
	//var err error
	result := db.Create(value)
	return result.Error, result.RowsAffected
}

func Save(db *gorm.DB, v interface{}) error {
	var err error
	err = db.Save(v).Error
	return err
}

// update语句 不允许使用orm特性
