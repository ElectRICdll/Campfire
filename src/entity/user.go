package entity

type ID uint

type User struct {
	ID        ID `gorm:"primaryKey"`
	Email     string
	Name      string
	AvatarUrl string
	Signature string
	Status    int
	Token     string
	IsOnline  bool `gorm:"-"`
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

type UserDTO struct {
	ID        ID     `json:"id,omitempty" uri:"user_id" binding:"required"`
	Email     string `json:"email,omitempty"`
	Name      string `json:"name,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Signature string `json:"signature,omitempty"`
	Status    int    `json:"status,omitempty"`
}
