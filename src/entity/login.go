package entity

type LoginDTO struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
	WS    string `json:"ws"`
}
