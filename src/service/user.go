package service

import (
	"campfire/cache"
	"campfire/dao"
	. "campfire/entity"
	. "campfire/util"
)

type UserService interface {
	UserInfo(id uint) (UserDTO, error)

	userInfo(id uint) (User, error)

	FindUsersByName(name string) ([]UserDTO, error)

	findUsersByName(name string) ([]User, error)

	EditUserInfo(user UserDTO) error

	ChangeEmail(userID uint, email string) error

	ChangePassword(userID uint, password string) error

	online(user *User)

	offline(id uint)

	Tasks(userID uint) ([]TaskDTO, error)

	PrivateCamps(userID uint) ([]CampDTO, error)

	PublicCamps(userID uint) ([]BriefCampDTO, error)

	Projects(userID uint) ([]BriefProjectDTO, error)
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
	err := s.userQuery.SetUserInfo(User{
		ID:    userID,
		Email: email,
	})
	// TODO
	return err
}

func (s *userService) ChangePassword(userID uint, password string) error {
	err := s.userQuery.SetUserInfo(User{ID: userID, Password: password})
	return err
}

func (s *userService) UserInfo(id uint) (UserDTO, error) {
	user, err := s.userQuery.UserInfoByID(id)

	return user.DTO(), err
}

func (s *userService) userInfo(id uint) (User, error) {
	user, err := s.userQuery.UserInfoByID(id)

	return user, err
}

func (s *userService) EditUserInfo(dto UserDTO) error {
	if err := s.userQuery.SetUserInfo(User{
		ID:        dto.ID,
		Name:      dto.Name,
		Signature: dto.Signature,
		AvatarUrl: dto.AvatarUrl,
	}); err != nil {
		return err
	}

	user := cache.TestUsers[dto.ID]
	if user == nil {
		user = &User{ID: dto.ID}
	}
	if len(dto.Name) != 0 {
		user.Name = dto.Name
	}
	if len(dto.Signature) != 0 {
		user.Signature = dto.Signature
	}
	if dto.Status != 0 {
		user.Status = dto.Status
	}

	return nil
}

func (s *userService) FindUsersByName(name string) ([]UserDTO, error) {
	users, err := s.userQuery.FindUsersByName(name)
	userDTOs := []UserDTO{}
	for _, user := range users {
		userDTOs = append(userDTOs, user.DTO())
	}

	return userDTOs, err
}

func (s *userService) findUsersByName(name string) ([]User, error) {
	users, err := s.userQuery.FindUsersByName(name)

	return users, ExternalError{
		Message: err.Error(),
	}
}

func (s *userService) Tasks(userID uint) ([]TaskDTO, error) {
	res, err := s.userQuery.TasksOfUser(userID)
	tasks := []TaskDTO{}
	for _, task := range res {
		tasks = append(tasks, TaskDTO{
			ID:      task.ID,
			OwnerID: task.OwnerID,
			//ReceiversID: task.Receiver,
			Title:   task.Title,
			Content: task.Content,
			BeginAt: task.BeginAt,
			EndAt:   task.EndAt,
			Status:  task.Status,
		})
	}
	return tasks, err
}

func (s *userService) PrivateCamps(userID uint) ([]CampDTO, error) {
	res, err := s.userQuery.PrivateCampsOfUser(userID)
	if err != nil {
		return nil, err
	}
	camps := []CampDTO{}
	for _, camp := range res {
		camps = append(camps, CampDTO{
			ID:   camp.ID,
			Name: camp.Name,
			//Members: camp.Members,
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
			OwnerID:      camp.OwnerID,
			MembersCount: len(camp.Members),
		})
	}
	return camps, err
}

func (s *userService) Projects(userID uint) ([]BriefProjectDTO, error) {
	res, err := s.userQuery.ProjectsOfUser(userID)
	projs := []BriefProjectDTO{}
	for _, proj := range res {
		projs = append(projs, BriefProjectDTO{
			ID:          proj.ID,
			Title:       proj.Title,
			Description: proj.Description,
		})
	}
	return projs, err
}

func (s *userService) online(user *User) {
	//s.onlineUsers[user.uint] = user
}

func (s *userService) offline(id uint) {
	//delete(s.onlineUsers, id)
}
