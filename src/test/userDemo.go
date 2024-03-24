package test

import (
	"campfire/entity"
	"campfire/log"
)

var (
	UserData = []entity.User{
		{
			Email:    "hare@email.com",
			Name:     "electric",
			Password: "420204",
		},
		{
			Email: "test1@email.com",
			Name: "test1",
			Password: "test1",
		},
	}
)

func UserDemo() {
	res, err := TestResource.Login(UserData[1].Email, UserData[1].Password)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Infof("%v", res)
	log.Info("登录测试成功！")
}
