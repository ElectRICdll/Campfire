package dao

import (
	. "campfire/entity"
	. "campfire/util"
	"gorm.io/gorm"
)

type UserDao interface {
	CheckIdentity(email string, password string) (uint, error)

	UserInfoByID(userID uint) (User, error)

	FindUsersByName(name string) ([]User, error)

	SetUserInfo(user User) error

	CreateUser(user User) error

	TasksOfUser(userID uint) ([]Task, error)

	CampsOfUser(userID uint) ([]Camp, error)

	PrivateCampsOfUser(userID uint) ([]Camp, error)

	ProjectsOfUser(userID uint) ([]Project, error)
}

func NewUserDao() UserDao {
	return userDao{}
}

type userDao struct{}

func (d userDao) SetUserInfo(user User) error {
	//result := DB.Exec("UPDATE user_info SET email = % s , name = %s , password = %s , signature = %s , avatar_url = %s WHERE user_id = %d", user.Email, user.Name, user.Password, user.Signature, user.AvatarUrl, user.ID)
	var result = DB.Save(&user)
	return result.Error
}

func (d userDao) CreateUser(user User) error {
	res := DB.Create(&user)
	return res.Error
}

func (d userDao) TasksOfUser(userID uint) ([]Task, error) {
	tasks := make([]Task, 0)
	result := DB.Preload("Project").Joins("JOIN members ON members.id = tasks.receivers").Where("members.user_id = ?", userID).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (d userDao) CampsOfUser(userID uint) ([]Camp, error) {
	camps := make([]Camp, 0)
	if err := DB.Preload("Members").
		Preload("Members.Camp").
		Where("members.user_id = ?", userID).
		Where("camps.is_private = ?", false).
		Find(&camps).Error; err != nil {
		return nil, err
	}

	return camps, nil
}

func (d userDao) PrivateCampsOfUser(userID uint) ([]Camp, error) {
	camps := make([]Camp, 0)
	if err := DB.Preload("Members").
		Preload("Members.Camp").
		Where("members.user_id = ?", userID).
		Where("camps.is_private = ?", true).
		Find(&camps).Error; err != nil {
		return nil, err
	}

	return camps, nil
}

func (d userDao) ProjectsOfUser(userID uint) ([]Project, error) {
	var projects []Project
	result := DB.Preload("Members").
		Preload("Camps").
		Preload("Tasks").
		Joins("JOIN project_members ON project_members.proj_id = projects.id").
		Where("project_members.user_id = ?", userID).
		Find(&projects)

	if result.Error != nil {
		return nil, result.Error
	}

	return projects, nil
}

func (d userDao) CheckIdentity(email string, password string) (uint, error) {
	var id uint
	//result := DB.Raw("SELECT ID FROM user_info WHERE email = %s AND password = %s", email, password).Scan(&id)
	var result = DB.Where("email = ? AND password = ?", email, password).Find(&id)

	if result.Error == gorm.ErrRecordNotFound {
		return id, ExternalError{}
	}
	if result.Error != nil {
		return id, result.Error
	}

	return id, nil
}

func (d userDao) UserInfoByID(userID uint) (User, error) {
	user := User{}
	res := DB.Where("id = ?", userID).First(&user)
	if res.Error == gorm.ErrRecordNotFound {
		return user, NewExternalError("No such user.")
	} else if res != nil {
		return user, res.Error
	}
	return user, nil
}

func (d userDao) FindUsersByName(name string) ([]User, error) {
	var user []User
	var result = DB.Where("name LIKE ?", "%"+name+"%").Find(&user)

	if result.Error == gorm.ErrRecordNotFound {
		return user, ExternalError{}
	}
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
