package service

import (
	"campfire/dao"
	"campfire/entity"
)

type LoginService interface {
	Login(email string, password string) (entity.LoginDTO, error)

	Register(email string, password string)
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

func (s *loginService) Login(email string, password string) (entity.LoginDTO, error) {
	password = s.sec.encryptPassword(password)

	if id, err := s.query.CheckIdentity(email, password); err == nil {
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
		} else {
			return entity.LoginDTO{}, err
		}
	}

	return entity.LoginDTO{}, entity.ExternalError{
		Message: "no such user",
	}
}

func (s *loginService) Logout() {}

func (s *loginService) Register(email string, password string) {}
