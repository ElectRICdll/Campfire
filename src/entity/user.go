package entity

import "time"

type User struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Email      string    `gorm:"not null;unique" json:"email"`
	Name       string    `gorm:"not null" json:"username"`
	Password   string    `gorm:"size:60;not null" json:"-"`
	AvatarUrl  string    `json:"-"`
	Signature  string    `json:"signature"`
	Status     int       `json:"status"`
	Token      string    `gorm:"-" json:"token"`
	LastOnline time.Time `json:"lastOnline"`
}

type BriefUserDTO struct {
	ID        uint   `json:"id" uri:"user_id"`
	Email     string `json:"email"`
	Name      string `json:"username"`
	AvatarUrl string `json:"avatarUrl"`
	Signature string `json:"signature"`
	Status    int    `json:"status"`
	Token     string `json:"token"`
}

func (d User) BriefDTO() BriefUserDTO {
	return BriefUserDTO{
		Email:     d.Email,
		ID:        d.ID,
		Name:      d.Name,
		AvatarUrl: d.AvatarUrl,
		Signature: d.Signature,
		Status:    d.Status,
		Token:     d.Token,
	}
}
