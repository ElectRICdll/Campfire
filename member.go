package entity

type Member struct {
	*User

	Nickname string
	Title    string
}

type MemberDTO struct {
	UserDTO
	NickName string `json:"nickname"`
	Title    string `json:"member_title"`
}
