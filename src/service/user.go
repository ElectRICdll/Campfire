package service

import (
	"campfire/cache"
	"campfire/dao"
	"campfire/entity"
)

type UserService interface {
	UserInfo(id int) (entity.UserDTO, error)

	userInfo(id int) (entity.User, error)

	FindUsersByName(name string) ([]entity.UserDTO, error)

	findUsersByName(name string) ([]entity.User, error)

	ChangeUserInfo(id string) error

	CampsiteList(campsiteId int) ([]entity.UserDTO, error)

	campsiteList(campsiteId int) ([]entity.User, error)

	online(user *entity.User)

	offline(id entity.ID)
}

func NewUserService() UserService {
	return &userService{
		dao.UserDaoContainer,
		cache.TestUsers,
	}
}

type userService struct {
	query       dao.UserDao
	onlineUsers map[entity.ID]*entity.User
}

func (s *userService) UserInfo(id int) (entity.UserDTO, error) {
	user, err := s.query.UserInfoByID(id)

	return user.DTO(), err
}

func (s *userService) userInfo(id int) (entity.User, error) {
	user, err := s.query.UserInfoByID(id)

	return user, err
}

func (s *userService) FindUsersByName(name string) ([]entity.UserDTO, error) {
	users, err := s.query.FindUsersByName(name)
	userDTOs := []entity.UserDTO{}
	for _, user := range users {
		userDTOs = append(userDTOs, user.DTO())
	}

	return userDTOs, entity.ExternalError{
		Message: err.Error(),
	}
}

func (s *userService) findUsersByName(name string) ([]entity.User, error) {
	users, err := s.query.FindUsersByName(name)

	return users, entity.ExternalError{
		Message: err.Error(),
	}
}

// TODO
func (s *userService) CampsiteList(campsiteId int) ([]entity.UserDTO, error) {
	return nil, nil
}

// TODO
func (s *userService) campsiteList(campsiteId int) ([]entity.User, error) {
	return nil, nil
}

// TODO
func (s *userService) ChangeUserInfo(id string) error {
	return nil
}

func (s *userService) online(user *entity.User) {
	s.onlineUsers[user.ID] = user
}

func (s *userService) offline(id entity.ID) {
	delete(s.onlineUsers, id)
}
