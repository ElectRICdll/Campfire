package service

import (
	"campfire/dao"
	. "campfire/entity"
	. "campfire/util"
)

type UserService interface {
	UserInfo(id uint) (User, error)

	userInfo(id uint) (User, error)

	FindUsersByName(name string) ([]User, error)

	findUsersByName(name string) ([]User, error)

	EditUserInfo(user User) error

	ChangeEmail(userID uint, email string) error

	ChangePassword(userID uint, password string) error

	online(user *User)

	offline(id uint)

	Tasks(userID uint) ([]Task, error)

	PrivateCamps(userID uint) ([]BriefCampDTO, error)

	PublicCamps(userID uint) ([]BriefCampDTO, error)

	Projects(userID uint) ([]Project, error)
}

func NewUserService() UserService {
	return &userService{
		dao.UserDaoContainer,
		dao.ProjectDaoContainer,
	}
}

type userService struct {
	userQuery dao.UserDao
	projQuery dao.ProjectDao
}

func (s *userService) ChangeEmail(userID uint, email string) error {
	if ok := ValidateEmail(email); !ok {
		return NewExternalError("illegal email format")
	}
	err := s.userQuery.SetUserInfo(User{
		ID:    userID,
		Email: email,
	})
	// TODO
	return err
}

func (s *userService) ChangePassword(userID uint, password string) error {
	if ok := ValidatePassword(password); !ok {
		return NewExternalError("illegal email format")
	}
	err := s.userQuery.SetUserInfo(User{ID: userID, Password: password})
	return err
}

func (s *userService) UserInfo(id uint) (User, error) {
	user, err := s.userQuery.UserInfoByID(id)
	if err != nil {
		return user, err
	}
	user.Avatar, err = FileToBase64(user.AvatarUrl)
	if err != nil {
		user.Avatar = ""
	}

	return user, nil
}

func (s *userService) userInfo(id uint) (User, error) {
	user, err := s.userQuery.UserInfoByID(id)

	return user, err
}

func (s *userService) EditUserInfo(dto User) error {
	if len(dto.Name) != 0 && !ValidateUsername(dto.Name) {
		return NewExternalError("illegal name format")
	}
	if err := s.userQuery.SetUserInfo(User{
		ID:        dto.ID,
		Name:      dto.Name,
		Signature: dto.Signature,
		AvatarUrl: dto.AvatarUrl,
	}); err != nil {
		return err
	}

	//user := cache.TestUsers[dto.ID]
	//if user == nil {
	//	user = &User{ID: dto.ID}
	//}
	//if len(dto.Name) != 0 {
	//	user.Name = dto.Name
	//}
	//if len(dto.Signature) != 0 {
	//	user.Signature = dto.Signature
	//}
	//if dto.Status != 0 {
	//	user.Status = dto.Status
	//}

	return nil
}

func (s *userService) FindUsersByName(name string) ([]User, error) {
	users, err := s.userQuery.FindUsersByName(name)

	return users, err
}

func (s *userService) findUsersByName(name string) ([]User, error) {
	users, err := s.userQuery.FindUsersByName(name)

	return users, ExternalError{
		Message: err.Error(),
	}
}

func (s *userService) Tasks(userID uint) ([]Task, error) {
	res, err := s.userQuery.TasksOfUser(userID)
	return res, err
}

func (s *userService) PrivateCamps(userID uint) ([]BriefCampDTO, error) {
	res, err := s.userQuery.PrivateCampsOfUser(userID)
	if err != nil {
		return nil, err
	}
	camps := []BriefCampDTO{}
	for _, camp := range res {
		camps = append(camps, BriefCampDTO{
			ID:        camp.ID,
			Name:      camp.Name,
			IsPrivate: camp.IsPrivate,
		})
	}
	return camps, nil
}

func (s *userService) PublicCamps(userID uint) ([]BriefCampDTO, error) {
	res, err := s.userQuery.PrivateCampsOfUser(userID)
	camps := []BriefCampDTO{}
	for _, camp := range res {
		camps = append(camps, BriefCampDTO{
			ID:           camp.ID,
			Name:         camp.Name,
			MembersCount: len(camp.Members) + 1,
			IsPrivate:    camp.IsPrivate,
		})
	}
	return camps, err
}

func (s *userService) Projects(userID uint) ([]Project, error) {
	res, err := s.userQuery.ProjectsOfUser(userID)
	projs := []Project{}
	for _, proj := range res {
		projs = append(projs, proj)
	}
	return projs, err
}

func (s *userService) online(user *User) {
	//s.onlineUsers[user.uint] = user
}

func (s *userService) offline(id uint) {
	//delete(s.onlineUsers, id)
}
