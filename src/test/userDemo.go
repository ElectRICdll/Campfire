package test

import (
	"campfire/entity"
	"campfire/log"
)

// UserData 这里是备用的测试用例，可以按喜好更改
var (
	UserData = []entity.User{
		{
			Email:    "hare@email.com",
			Name:     "electric",
			Password: "420204",
		},
		{
			Email:    "test1@email.com",
			Name:     "test1",
			Password: "test1",
		},
	}
)

// UserDemo 是运行UserService和LoginService的Demo方法
func UserDemo() {
	// 注册测试示例
	err := ResourceTest.Register(UserData[1].DTO(), UserData[1].Password)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Info("注册测试成功！")
	// 登录测试示例
	res, err := ResourceTest.Login(UserData[1].Email, UserData[1].Password)
	if err != nil {
		log.Error(err.Error())
		return
	}
	// 这个写法可以直接打印结构体
	log.Infof("%#v", res)
	log.Info("登录测试成功！")
}
