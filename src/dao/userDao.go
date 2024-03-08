package dao

import (
	"campfire/cache"
	"campfire/entity"
)

type UserDao interface {
	// CheckIdentity 用于登录接口的方法，向sql后台验证用户是否存在，返回其id即可。
	CheckIdentity(email string, password string) (int, error)

	UserInfoByID(userID int) (entity.User, error)

	FindUsersByName(name string) ([]entity.User, error)

	SetUserInfo(userID int, user entity.UserDTO) error

	SetPassword(userID int, userid int, password string) error

	CreateUser(user entity.User, password string) error
}

// func NewUserDaoTest() UserDao {
// 	return userDaoTest{}
// }

type userDaoTest struct{}

func (d userDaoTest) SetUserSign(userID int, signature string) error {
	result := db.Exec("UPDATE user_info SET signature = %s WHERE user_id = %d", signature, userID)
	if result.Error != nil {
		return result.Error
	}
	if result != nil {
		return nil
	}
	return entity.ExternalError{}
}

func (d userDaoTest) ChangePassword(userID int, p string) error {
	result := db.Exec("UPDATE user_info SET password = %s WHERE user_id = %d", p, userID)
	if result.Error != nil {
		return result.Error
	}
	if result != nil {
		return nil
	}
	return entity.ExternalError{}
}

func (d userDaoTest) ChangeEmail(userID int, email string) error {
	result := db.Exec("UPDATE user_info SET email = %s WHERE user_id = %d", email, userID)
	if result.Error != nil {
		return result.Error
	}
	if result != nil {
		return nil
	}
	return entity.ExternalError{}
}

func (d userDaoTest) SetUserName(userID int, name string) error {
	result := db.Exec("UPDATE user_info SET name = %s WHERE user_id = %d", name, userID)
	if result.Error != nil {
		return result.Error
	}
	if result != nil {
		return nil
	}
	return entity.ExternalError{}
}

func (d userDaoTest) CheckIdentity(email string, password string) (int, error) {
	var id int
	result := db.Raw("SELECT user_id FROM user_info WHERE email = %s AND password = %s", email, password).Scan(&id)
	if result.Error != nil {
		return 0, result.Error
	}
	if result != nil {
		return id, nil
	}
	return 0, entity.ExternalError{}
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
