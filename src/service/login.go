package service

import (
	"campfire/cache"
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
		NewUserService(),
		NewSecurityService(),
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
	user, err := s.query.CheckIdentity(email, password)
	if err != nil {
		return entity.LoginDTO{}, err
	}

	token, err := s.sec.tokenGenerate(user)
	if err != nil {
		return entity.LoginDTO{}, err
	}
	user.Token = token
	s.user.online(&user)
	cache.StoreUserInCache(user)
	return entity.LoginDTO{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
	}, nil
}

func (s *loginService) Register(dto entity.UserDTO, password string) error {
	p, err := s.sec.encryptPassword(password)
	if err != nil {
		return err
	}
	if err := s.query.CreateUser(entity.User{
		Email:    dto.Email,
		Name:     dto.Name,
		Password: p,
	}); err != nil {
		return err
	}
	return nil
}
