package entity

type LoginDTO struct {
	ID    uint   `json:"id"`
	Token string `json:"token"`
	WS    string `json:"ws"`
}
