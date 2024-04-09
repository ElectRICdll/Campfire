package service

import (
	"campfire/auth"
	"campfire/cache"
	"campfire/dao"
	"campfire/entity"
	"campfire/util"
)

type LoginService interface {
	Login(email string, password string) (entity.LoginDTO, error)

	Register(user entity.User, password string) (uint, error)

	EmailVerify(vefiryCode string) error
}

func NewLoginService() LoginService {
	return &loginService{
		dao.UserDaoContainer,
		NewUserService(),
		auth.SecurityInstance,
	}
}

type loginService struct {
	query dao.UserDao
	user  UserService
	sec   auth.SecurityGuard
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

	token, err := s.sec.TokenGenerate(user)
	if err != nil {
		return entity.LoginDTO{}, err
	}
	user.Token = token
	s.user.online(&user)
	cache.StoreUserInCache(user)
	return entity.LoginDTO{
		ID:         user.ID,
		Name:       user.Name,
		Token:      token,
		LastOnline: user.LastOnline,
	}, nil
}

func (s *loginService) Register(dto entity.User, password string) (uint, error) {
	if ok := util.ValidateUsername(dto.Name); !ok {
		return 0, util.NewExternalError("illegal username format")
	}

	if ok := util.ValidatePassword(password); !ok {
		return 0, util.NewExternalError("illegal password format")
	}

	p, err := s.sec.EncryptPassword(password)
	if err != nil {
		return 0, err
	}
	res, err := s.query.CreateUser(entity.User{
		Email:    dto.Email,
		Name:     dto.Name,
		Password: p,
	})
	if err != nil {
		return 0, err
	}
	return res, err
}
