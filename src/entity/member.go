package entity

type Member struct {
	*User

	Nickname string
	Title    string
}
