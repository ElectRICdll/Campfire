package service

import (
	"campfire/dao"
	"campfire/entity"
)

type LoginService interface {
	Login(email string, password string) (entity.LoginDTO, error)

	Register(user entity.UserDTO, password string) error

	EmailVerify(vefiryCode string) error
}

func NewLoginService() LoginService {
	return &loginService{
		dao.UserDaoContainer,
		UserServiceContainer,
		SecurityServiceContainer,
	}
}

type loginService struct {
	query dao.UserDao
	user  UserService
	sec   SecurityService
}

func (s *loginService) EmailVerify(vefiryCode string) error {
	//TODO implement me
	panic("implement me")
}

func (s *loginService) Login(email string, password string) (entity.LoginDTO, error) {
	password = s.sec.encryptPassword(password)

	id, err := s.query.CheckIdentity(email, password)
	if err != nil {
		return entity.LoginDTO{}, err
	}
	user, err := s.user.userInfo(id)
	if err != nil {
		return entity.LoginDTO{}, err
	}

	token, err := s.sec.tokenGenerate(user)
	if err == nil {
		user.Token = token
		s.user.online(&user)
		return entity.LoginDTO{
			ID:    id,
			Token: token,
		}, err
	}
	return entity.LoginDTO{}, err
}

func (s *loginService) Register(dto entity.UserDTO, password string) error {
	if err := s.query.CreateUser(entity.User{
		Email:    dto.Email,
		Name:     dto.Name,
		Password: s.sec.encryptPassword(password),
	}); err != nil {
		return err
	}
	return nil
}
