package dao

import (
	"campfire/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	UserDaoContainer    UserDao    = NewUserDao()
	ProjectDaoContainer ProjectDao = NewProjectDao()
	CampDaoContainer    CampDao    = NewCampDao()
	MessageDaoContainer MessageDao = NewMessageDao()
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
