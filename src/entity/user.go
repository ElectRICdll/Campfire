package entity

type User struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Email     string
	Name      string
	Password  string
	AvatarUrl string
	Signature string
	Status    int
	Token     string `gorm:"-"`
	IsOnline  bool   `gorm:"-"`
}

type UserDTO struct {
	ID        uint   `json:"id,omitempty" uri:"user_id"`
	Email     string `json:"email,omitempty"`
	Name      string `json:"name,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Signature string `json:"signature,omitempty"`
	Status    int    `json:"status,omitempty"`
}

func (d User) DTO() UserDTO {
	return UserDTO{
		Email:     d.Email,
		ID:        d.ID,
		Name:      d.Name,
		AvatarUrl: d.AvatarUrl,
		Signature: d.Signature,
		Status:    d.Status,
	}
}

func UsersDTO(users []User) []UserDTO {
	res := []UserDTO{}
	for _, user := range users {
		res = append(res, user.DTO())
	}
	return res
}
