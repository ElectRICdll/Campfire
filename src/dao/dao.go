package dao

import (
	"campfire/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	UserDaoContainer    UserDao    = nil
	ProjectDaoContainer ProjectDao = nil
)

// 需要给让db连接数据库
var db *gorm.DB

//--------------------

func DBConn() *gorm.DB {
	db, err := gorm.Open(mysql.Open(util.CONFIG.SQLConn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}
