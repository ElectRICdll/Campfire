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

	EditUserInfo(user entity.UserDTO) error

	ChangePassword(userId int, password string) error

	online(user *entity.User)

	offline(id entity.ID)

	Tents(userId int) ([]entity.BriefTentDTO, error)

	Projects(userId int) ([]entity.BriefProjectDTO, error)

	Campsites(userId int) ([]entity.BriefCampDTO, error)
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

func (s *userService) Tents(userId int) ([]entity.BriefTentDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) Projects(userId int) ([]entity.BriefProjectDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) Campsites(userId int) ([]entity.BriefCampDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) ChangePassword(userId int, password string) error {
	err := s.query.SetPassword(userId, password)
	return err
}

func (s *userService) UserInfo(id int) (entity.UserDTO, error) {
	user, err := s.query.UserInfoByID(id)

	return user.DTO(), err
}

func (s *userService) userInfo(id int) (entity.User, error) {
	user, err := s.query.UserInfoByID(id)

	return user, err
}

func (s *userService) EditUserInfo(dto entity.UserDTO) error {
	if err := s.query.SetUserInfo(dto); err != nil {
		return err
	}

	user := cache.TestUsers[(entity.ID)(dto.ID)]
	if user == nil {
		user = &entity.User{ID: (entity.ID)(dto.ID)}
	}
	if len(dto.Name) != 0 {
		user.Name = dto.Name
	}
	if len(dto.Signature) != 0 {
		user.Signature = dto.Signature
	}
	if dto.Status != 0 {
		user.Status = (entity.Status)(dto.Status)
	}

	return nil
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

func (s *userService) online(user *entity.User) {
	s.onlineUsers[user.ID] = user
}

func (s *userService) offline(id entity.ID) {
	delete(s.onlineUsers, id)
}
