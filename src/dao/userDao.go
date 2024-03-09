package dao

import (
	"campfire/cache"
	. "campfire/entity"
	"gorm.io/gorm"
)

type UserDao interface {
	// CheckIdentity 用于登录接口的方法，向sql后台验证用户是否存在，返回其id即可。
	CheckIdentity(email string, password string) (uint, error)

	UserInfoByID(userID uint) (User, error)

	FindUsersByName(name string) ([]User, error)

	SetUserInfo(user User) error

	SetPassword(userID uint, password string) error

	CreateUser(user User) error

	TasksOfUser(userID uint) ([]Task, error)

	CampsOfUser(userID uint) ([]Camp, error)

	PrivateCampsOfUser(userID uint) ([]Camp, error)

	ProjectsOfUser(userID uint) ([]Project, error)
}

// func NewUserDaoTest() UserDao {
// 	return userDao{}
// }

type userDao struct{}

func (d userDao) SetUserSign(userID uint, signature string) error {
	result := DB.Exec("UPDATE user_info SET signature = %s WHERE user_id = %d", signature, userID)
	if result.Error != nil {
		return result.Error
	}
	if result != nil {
		return nil
	}
	return ExternalError{}
}

func (d userDao) ChangePassword(userID uint, p string) error {
	result := DB.Exec("UPDATE user_info SET password = %s WHERE user_id = %d", p, userID)
	if result.Error != nil {
		return result.Error
	}
	if result != nil {
		return nil
	}
	return ExternalError{}
}

func (d userDao) ChangeEmail(userID uint, email string) error {
	result := DB.Exec("UPDATE user_info SET email = %s WHERE user_id = %d", email, userID)
	if result.Error != nil {
		return result.Error
	}
	if result != nil {
		return nil
	}
	return ExternalError{}
}

func (d userDao) SetUserName(userID uint, name string) error {
	result := DB.Exec("UPDATE user_info SET name = %s WHERE user_id = %d", name, userID)
	if result.Error != nil {
		return result.Error
	}
	if result != nil {
		return nil
	}
	return ExternalError{}
}

func (d userDao) CheckIdentity(email string, password string) (uint, error) {
	var id uint
	result := DB.Raw("SELECT user_id FROM user_info WHERE email = %s AND password = %s", email, password).Scan(&id)
	if result.Error != nil {
		return 0, result.Error
	}
	if result != nil {
		return id, nil
	}
	return 0, ExternalError{}
}

func (d userDao) UserInfoByID(id uint) (User, error) {
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

func (d userDao) FindUsersByName(name string) ([]User, error) {
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
func (d userDao) SetAvatar(id uint, url string) error {
	return nil
}

type userDaoTest struct{}

func NewUserDaoTest() UserDao {
	return userDaoTest{}
}

func (s userDaoTest) CheckIdentity(email string, password string) (uint, error) {
	user := User{
		Email:    email,
		Password: password,
	}
	res := DB.Where("email = ?", email).First(&user)
	if res.Error == gorm.ErrRecordNotFound {
		return 0, NewExternalError("No such user.")
	} else if user.Password != password {
		return 0, NewExternalError("Wrong password or account.")
	} else if res.Error != nil {
		return 0, res.Error
	}
	return user.ID, nil
}

func (s userDaoTest) UserInfoByID(userID uint) (User, error) {
	user := User{}
	res := DB.Where("id = ?", userID).First(&user)
	if res.Error == gorm.ErrRecordNotFound {
		return user, NewExternalError("No such user.")
	} else if res != nil {
		return user, res.Error
	}
	return user, nil
}

func (s userDaoTest) FindUsersByName(name string) ([]User, error) {
	//TODO implement me
	panic("implement me")
}

func (s userDaoTest) SetUserInfo(user User) error {
	//TODO implement me
	panic("implement me")
}

func (s userDaoTest) SetPassword(userID uint, password string) error {
	//TODO implement me
	panic("implement me")
}

func (s userDaoTest) CreateUser(user User) error {
	res := DB.Create(&user)
	return res.Error
}

func (s userDaoTest) TasksOfUser(userID uint) ([]Task, error) {
	//TODO implement me
	panic("implement me")
}

func (s userDaoTest) CampsOfUser(userID uint) ([]Camp, error) {
	//TODO implement me
	panic("implement me")
}

func (s userDaoTest) PrivateCampsOfUser(userID uint) ([]Camp, error) {
	//TODO implement me
	panic("implement me")
}

func (s userDaoTest) ProjectsOfUser(userID uint) ([]Project, error) {
	//TODO implement me
	panic("implement me")
}
