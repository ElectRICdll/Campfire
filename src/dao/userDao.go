package dao

import (
	"campfire/cache"
	. "campfire/entity"
	. "campfire/util"
	"golang.org/x/crypto/bcrypt"

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
	//result := DB.Exec("UPDATE user_info SET signature = %s WHERE ID = %d", signature, userID)
	var user User
	var result = DB.First(&user, userID)

	if result.Error != nil {
		return result.Error
	}

	user.Signature = signature
	result = DB.Save(&user)

	if result.Error != nil {
		return result.Error
	}

	if result == nil {
		return ExternalError{}
	}

	return nil
}

func (d userDao) ChangePassword(userID uint, p string) error {
	//result := DB.Exec("UPDATE user_info SET password = %s WHERE ID = %d", p, userID)
	var user User
	var result = DB.First(&user, userID)

	if result.Error != nil {
		return result.Error
	}

	user.Password = p
	result = DB.Save(&user)

	if result.Error != nil {
		return result.Error
	}

	if result == nil {
		return ExternalError{}
	}

	return nil
}

func (d userDao) ChangeEmail(userID uint, email string) error {
	//result := DB.Exec("UPDATE user_info SET email = %s WHERE ID = %d", email, userID)
	var user User
	var result = DB.First(&user, userID)

	if result.Error != nil {
		return result.Error
	}

	user.Email = email
	result = DB.Save(&user)

	if result.Error != nil {
		return result.Error
	}

	if result == nil {
		return ExternalError{}
	}

	return nil
}

func (d userDao) SetUserName(userID uint, name string) error {
	//result := DB.Exec("UPDATE user_info SET name = %s WHERE ID = %d", name, userID)
	var user User
	var result = DB.First(&user, userID)

	if result.Error != nil {
		return result.Error
	}

	user.Name = name
	result = DB.Save(&user)

	if result.Error != nil {
		return result.Error
	}

	if result == nil {
		return ExternalError{}
	}

	return nil
}

func (d userDao) CheckIdentity(email string, password string) (uint, error) {
	var id uint
	//result := DB.Raw("SELECT ID FROM user_info WHERE email = %s AND password = %s", email, password).Scan(&id)
	var result = DB.Where("email = ? and password = ?", email, password).Find(&id)

	if result.Error != nil {
		return id, result.Error
	}

	if result == nil {
		return id, ExternalError{}
	}

	return id, nil
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
	user := User{}
	res := DB.Where("email = ?", email).First(&user)
	if res.Error == gorm.ErrRecordNotFound {
		return 0, NewExternalError("No such user.")
	} else if err := bcrypt.CompareHashAndPassword(([]byte)(user.Password), ([]byte)(password)); err != nil {
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
	var user []User
	var result = DB.Where("name = ?", name).Find(&user)

	if result.Error != nil {
		return user, result.Error
	}

	if result == nil {
		return user, ExternalError{}
	}

	return user, nil
}

func (s userDaoTest) SetUserInfo(user User) error {
	//result := DB.Exec("UPDATE user_info SET email = % s , name = %s , password = %s , signature = %s , avatar_url = %s WHERE user_id = %d", user.Email, user.Name, user.Password, user.Signature, user.AvatarUrl, user.ID)
	var result = DB.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	if result != nil {
		return nil
	}
	return ExternalError{}
}

func (s userDaoTest) SetPassword(userID uint, password string) error {
	var user User
	var result = DB.First(&user, userID)

	if result.Error != nil {
		return result.Error
	}

	user.Password = password
	result = DB.Save(&user)

	if result.Error != nil {
		return result.Error
	}

	if result == nil {
		return ExternalError{}
	}

	return nil
}

func (s userDaoTest) CreateUser(user User) error {
	res := DB.Create(&user)
	return res.Error
}

func (s userDaoTest) TasksOfUser(userID uint) ([]Task, error) {
	var task []Task
	var result = DB.Where("OwnerID = ?", userID).Find(&task)

	if result.Error != nil {
		return task, result.Error
	}

	if result == nil {
		return task, ExternalError{}
	}

	return task, nil
}

func (s userDaoTest) CampsOfUser(userID uint) ([]Camp, error) {
	var camp []Camp
	var result = DB.Where("OwnerID = ?", userID).Find(&camp)

	if result.Error != nil {
		return camp, result.Error
	}

	if result == nil {
		return camp, ExternalError{}
	}

	return camp, nil
}

func (s userDaoTest) PrivateCampsOfUser(userID uint) ([]Camp, error) {
	//待改表
	return nil, nil
	// var camp []Camp
	// result := DB.Raw("SELECT * FROM camp WHERE leader = %d and isprivate = 1", userID).Scan(&camp)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }
	// if result != nil {
	// 	return camp, nil
	// }
	// return nil, ExternalError{}
}

func (s userDaoTest) ProjectsOfUser(userID uint) ([]Project, error) {
	var project []Project
	var result = DB.Where("OwnerID = ?", userID).Find(&project)

	if result.Error != nil {
		return project, result.Error
	}

	if result == nil {
		return project, ExternalError{}
	}

	return project, nil
}
