package main

import "campfire/service"

var TestResource = struct {
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

func TestDemo() {
}
