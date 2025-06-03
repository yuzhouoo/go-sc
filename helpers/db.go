package helpers

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"go-session-demo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	conn *gorm.DB
}

var db *gorm.DB

func (d *DB) InitDB() (dbConn *gorm.DB, err error) {
	dsn := "root:root@tcp(192.168.11.11:13306)/go-test?charset=utf8mb4&parseTime=True&loc=Local"

	if db == nil {
		conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, errors.New("数据库初始化失败")
		}

		db = conn
		d.handleAutoMigrate()

		return conn, nil
	}

	return db, nil

}

/*
*自动更新表
 */
func (d *DB) handleAutoMigrate() {
	err := db.AutoMigrate(&models.AccountModel{})
	if err != nil {
		return
	}
}
