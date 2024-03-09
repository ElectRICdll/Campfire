package dao

import (
	"campfire/cache"
	. "campfire/entity"
)

type UserDao interface {
	// CheckIdentity 用于登录接口的方法，向sql后台验证用户是否存在，返回其id即可。
	CheckIdentity(email string, password string) (uint, error)

	UserInfoByID(userID uint) (User, error)

	FindUsersByName(name string) ([]User, error)

	SetUserInfo(user User) error

	SetPassword(userID uint, password string) error

	CreateUser(user User, password string) error

	TasksOfUser(userID uint) ([]Task, error)

	CampsOfUser(userID uint) ([]Camp, error)

	PrivateCampsOfUser(userID uint) ([]Camp, error)

	ProjectsOfUser(userID uint) ([]Project, error)
}

// func NewUserDaoTest() UserDao {
// 	return userDaoTest{}
// }

type userDaoTest struct{}

func (d userDaoTest) SetUserSign(userID uint, signature string) error {
	result := db.Exec("UPDATE user_info SET signature = %s WHERE user_id = %d", signature, userID)
	if result.Error != nil {
		return result.Error
	}
	if result != nil {
		return nil
	}
	return ExternalError{}
}

func (d userDaoTest) ChangePassword(userID uint, p string) error {
	result := db.Exec("UPDATE user_info SET password = %s WHERE user_id = %d", p, userID)
	if result.Error != nil {
		return result.Error
	}
	if result != nil {
		return nil
	}
	return ExternalError{}
}

func (d userDaoTest) ChangeEmail(userID uint, email string) error {
	result := db.Exec("UPDATE user_info SET email = %s WHERE user_id = %d", email, userID)
	if result.Error != nil {
		return result.Error
	}
	if result != nil {
		return nil
	}
	return ExternalError{}
}

func (d userDaoTest) SetUserName(userID uint, name string) error {
	result := db.Exec("UPDATE user_info SET name = %s WHERE user_id = %d", name, userID)
	if result.Error != nil {
		return result.Error
	}
	if result != nil {
		return nil
	}
	return ExternalError{}
}

func (d userDaoTest) CheckIdentity(email string, password string) (uint, error) {
	var id uint
	result := db.Raw("SELECT user_id FROM user_info WHERE email = %s AND password = %s", email, password).Scan(&id)
	if result.Error != nil {
		return 0, result.Error
	}
	if result != nil {
		return id, nil
	}
	return 0, ExternalError{}
}

func (d userDaoTest) UserInfoByID(id uint) (User, error) {
	if id == 1 {
		return *cache.TestUsers[1], nil
	} else {
		return User{
			ID:        65535,
			Email:     "",
			Name:      "else",
			AvatarUrl: "",
			Signature: "something else",
		}, nil
	}
}

func (d userDaoTest) FindUsersByName(name string) ([]User, error) {
	return []User{
		{
			ID:        1,
			Email:     "hare@mail.com",
			Name:      "electric",
			AvatarUrl: "",
			Signature: "",
		},
	}, nil
}

// TODO
func (d userDaoTest) SetAvatar(id uint, url string) error {
	return nil
}
