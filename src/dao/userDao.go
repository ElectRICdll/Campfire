package dao

import (
	. "campfire/entity"
	. "campfire/util"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserDao interface {
	CheckIdentity(email string, password string) (User, error)

	UserInfoByID(userID uint) (User, error)

	FindUsersByName(name string) ([]User, error)

	SetUserInfo(user User) error

	CreateUser(user User) (uint, error)

	TasksOfUser(userID uint) ([]Task, error)

	CampsOfUser(userID uint) ([]Camp, error)

	PrivateCampsOfUser(userID uint) ([]Camp, error)

	ProjectsOfUser(userID uint) ([]Project, error)
}

func NewUserDao() UserDao {
	return userDao{
		db: DBConn(),
	}
}

type userDao struct {
	db *gorm.DB
}

func (d userDao) SetUserInfo(user User) error {
	//result := d.db.Exec("UPDATE user_info SET email = % s , name = %s , password = %s , signature = %s , avatar_url = %s WHERE user_id = %d", user.Email, user.Name, user.Password, user.Signature, user.AvatarUrl, user.ID)
	var result = d.db.Updates(&user)
	return result.Error
}

func (d userDao) CreateUser(user User) (uint, error) {
	res := d.db.Create(&user)
	if res.Error == gorm.ErrDuplicatedKey {
		return 0, NewExternalError("The email has already been registered")
	}
	if res.Error != nil {
		return 0, res.Error
	}

	return user.ID, nil
}

func (d userDao) TasksOfUser(userID uint) ([]Task, error) {
	tasks := make([]Task, 0)
	result := d.db.Preload("Project").Joins("JOIN members ON members.id = tasks.receivers").Where("members.user_id = ?", userID).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (d userDao) CampsOfUser(userID uint) ([]Camp, error) {
	var camps []Camp

	var members []Member
	if err := d.db.Where("user_id = ?", userID).Find(&members).Error; err != nil {
		return nil, err
	}

	var campIDs []uint
	for _, member := range members {
		campIDs = append(campIDs, member.CampID)
	}

	if err := d.db.Where("id IN (?) and is_private = 0", campIDs).Find(&camps).Error; err != nil {
		return nil, err
	}

	return camps, nil
}

func (d userDao) PrivateCampsOfUser(userID uint) ([]Camp, error) {
	var camps []Camp

	var members []Member
	if err := d.db.Where("user_id = ?", userID).Find(&members).Error; err != nil {
		return nil, err
	}

	var campIDs []uint
	for _, member := range members {
		campIDs = append(campIDs, member.CampID)
	}

	if err := d.db.Preload("Members.User").Where("id IN (?) and is_private = 1", campIDs).Find(&camps).Error; err != nil {
		return nil, err
	}

	return camps, nil
}

func (d userDao) ProjectsOfUser(userID uint) ([]Project, error) {
	var projects []Project

	var members []ProjectMember
	if err := d.db.Where("user_id = ?", userID).Find(&members).Error; err != nil {
		return nil, err
	}

	var projectIDs []uint
	for _, member := range members {
		projectIDs = append(projectIDs, member.ProjID)
	}

	if err := d.db.Where("id IN (?)", projectIDs).Preload("Owner").Preload("Members.User").Find(&projects).Error; err != nil {
		return nil, err
	}

	return projects, nil
}

func (d userDao) CheckIdentity(email string, password string) (User, error) {
	user := User{}
	//result := d.db.Raw("SELECT ID FROM user_info WHERE email = %s AND password = %s", email, password).Scan(&id)
	var result = d.db.Where("email = ?", email).Find(&user)

	if err := bcrypt.CompareHashAndPassword(([]byte)(user.Password), ([]byte)(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return User{}, NewExternalError("wrong password")
		}
		return User{}, err
	}

	if result.Error == gorm.ErrRecordNotFound {
		return User{}, NewExternalError("No such user.")
	}
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func (d userDao) UserInfoByID(userID uint) (User, error) {
	user := User{}
	res := d.db.Where("id = ?", userID).First(&user)
	if res.Error == gorm.ErrRecordNotFound {
		return user, NewExternalError("No such user.")
	} else if res != nil {
		return user, res.Error
	}
	return user, nil
}

func (d userDao) FindUsersByName(name string) ([]User, error) {
	var user []User
	var result = d.db.Where("name LIKE ?", "%"+name+"%").Find(&user)

	if result.Error == gorm.ErrRecordNotFound {
		return user, ExternalError{}
	}
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
