package dao

import (
	"campfire/cache"
	"campfire/entity"
)

type UserDao interface {
	// CheckIdentity 用于登录接口的方法，向sql后台验证用户是否存在，返回其id即可。
	CheckIdentity(email string, password string) (int, error)

	UserInfoByID(userId int) (entity.User, error)

	FindUsersByName(name string) ([]entity.User, error)

	SetUserName(userId int, name string) error

	SetUserSign(userId int, signature string) error

	// SetAvatar 设置头像
	SetAvatar(userId int, url string) error

	ChangePassword(userId int, p string) error

	ChangeEmail(userId int, email string) error
}

func NewUserDaoTest() UserDao {
	return userDaoTest{}
}

type userDaoTest struct{}

func (d userDaoTest) SetUserSign(userId int, signature string) error {
	//TODO implement me
	panic("implement me")
}

func (d userDaoTest) ChangePassword(userId int, p string) error {
	//TODO implement me
	panic("implement me")
}

func (d userDaoTest) ChangeEmail(userId int, email string) error {
	//TODO implement me
	panic("implement me")
}

func (d userDaoTest) SetUserName(userId int, name string) error {
	//TODO implement me
	panic("implement me")
}

func (d userDaoTest) CheckIdentity(email string, password string) (int, error) {
	return 1, nil
}

func (d userDaoTest) UserInfoByID(id int) (entity.User, error) {
	if id == 1 {
		return *cache.TestUsers[1], nil
	} else {
		return entity.User{
			ID:        65535,
			Email:     "",
			Name:      "else",
			Avatar:    "",
			Signature: "something else",
		}, nil
	}
}

func (d userDaoTest) FindUsersByName(name string) ([]entity.User, error) {
	return []entity.User{
		{
			ID:        1,
			Email:     "hare@mail.com",
			Name:      "electric",
			Avatar:    "",
			Signature: "",
		},
	}, nil
}

// TODO
func (d userDaoTest) SetAvatar(id int, url string) error {
	return nil
}
