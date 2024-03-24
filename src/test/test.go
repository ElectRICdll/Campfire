package test

import (
	"campfire/service"
)

// ResourceTest 负责提供依赖注入，用这个类可以调用到所有的业务层方法
var ResourceTest = struct {
	service.CampService
	service.ProjectService
	service.LoginService
	service.UserService
	service.MessageService
	service.TaskService
}{
	service.CampServiceContainer,
	service.ProjectServiceContainer,
	service.LoginServiceContainer,
	service.UserServiceContainer,
	service.MessageServiceContainer,
	service.TaskServiceContainer,
}

// Demo 将直接被主函数调用
func Demo() {
	// UserDemo是专门调试User业务的Demo方法
	UserDemo()
}
