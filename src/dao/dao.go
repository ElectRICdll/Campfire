package dao

import (
	"campfire/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	UserDaoContainer    UserDao    = NewUserDaoTest()
	ProjectDaoContainer ProjectDao = nil
	CampDaoContainer    CampDao    = nil
	MessageDaoContainer MessageDao = nil
)

var DB *gorm.DB = DBConn()

//--------------------

func DBConn() *gorm.DB {
	db, err := gorm.Open(mysql.Open(util.CONFIG.SQLConn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}
