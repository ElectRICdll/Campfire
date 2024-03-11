package entity

type LoginDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"username"`
	Token     string `json:"token"`
	AvatarUrl string `json:"avatar_url"`
}
