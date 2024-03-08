package service

import (
	"campfire/cache"
	"campfire/dao"
	. "campfire/entity"
)

type UserService interface {
	UserInfo(id ID) (UserDTO, error)

	userInfo(id ID) (User, error)

	FindUsersByName(name string) ([]UserDTO, error)

	findUsersByName(name string) ([]User, error)

	EditUserInfo(user UserDTO) error

	ChangePassword(userID ID, password string) error

	online(user *User)

	offline(id ID)

	Tasks(userID ID) ([]TaskDTO, error)

	PrivateCamps(userID ID) ([]CampDTO, error)

	PublicCamps(userID ID) ([]BriefCampDTO, error)

	Projects(userID ID) ([]BriefProjectDTO, error)
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

func (s *userService) ChangePassword(userID ID, password string) error {
	err := s.userQuery.SetPassword(userID, password)
	return err
}

func (s *userService) UserInfo(id ID) (UserDTO, error) {
	user, err := s.userQuery.UserInfoByID(id)

	return user.DTO(), err
}

func (s *userService) userInfo(id ID) (User, error) {
	user, err := s.userQuery.UserInfoByID(id)

	return user, err
}

func (s *userService) EditUserInfo(dto UserDTO) error {
	if err := s.userQuery.SetUserInfo(User{
		ID:        dto.ID,
		Name:      dto.Name,
		Email:     dto.Email,
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

	return userDTOs, ExternalError{
		Message: err.Error(),
	}
}

func (s *userService) findUsersByName(name string) ([]User, error) {
	users, err := s.userQuery.FindUsersByName(name)

	return users, ExternalError{
		Message: err.Error(),
	}
}

func (s *userService) Tasks(userID ID) ([]TaskDTO, error) {
	res, err := s.userQuery.TasksOfUser(userID)
	tasks := []TaskDTO{}
	for _, task := range res {
		tasks = append(tasks, TaskDTO{
			ID:         task.ID,
			OwnerID:    task.OwnerID,
			ReceiverID: task.ReceiverID,
			Title:      task.Title,
			Content:    task.Content,
			Begin:      task.Begin,
			End:        task.End,
			Status:     task.Status,
		})
	}
	return tasks, err
}

func (s *userService) PrivateCamps(userID ID) ([]CampDTO, error) {
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

func (s *userService) PublicCamps(userID ID) ([]BriefCampDTO, error) {
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

func (s *userService) Projects(userID ID) ([]BriefProjectDTO, error) {
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
	//s.onlineUsers[user.ID] = user
}

func (s *userService) offline(id ID) {
	//delete(s.onlineUsers, id)
}
