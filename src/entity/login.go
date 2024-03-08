package entity

type LoginDTO struct {
	ID    ID     `json:"id"`
	Token string `json:"token"`
	WS    string `json:"ws"`
}
