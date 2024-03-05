package entity

type ID int
type Status int

type User struct {
	ID        ID
	Email     string
	Name      string
	Avatar    string
	Signature string
	Status    Status
	Token     string
	IsOnline  bool
}

func (d User) DTO() UserDTO {
	return UserDTO{
		Email:     d.Email,
		ID:        (int)(d.ID),
		Name:      d.Name,
		AvatarUrl: d.Avatar,
		Signature: d.Signature,
		Status:    (int)(d.Status),
	}
}
