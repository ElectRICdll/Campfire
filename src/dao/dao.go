package dao

import "gorm.io/gorm"

var (
	UserDaoContainer UserDao = nil
)

//需要给让db连接数据库
var db *gorm.DB

//--------------------

type DB struct {
}
