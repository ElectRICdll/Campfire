package entity

type MemberList map[int]*Member

type Member struct {
	ID     ID
	UserID ID
	ProjID ID
	CampID ID

	Nickname string
	Title    string

	User User    `gorm:"foreignKey:UserID"`
	Proj Project `gorm:"foreignKey:ProjID"`
	Camp Camp    `gorm:"foreignKey:CampID"`
}

type MemberDTO struct {
	ID        ID     `json:"id,omitempty" uri:"user_id" binding:"required"`
	UserID    ID     `json:"u_id,omitempty"`
	ProjID    ID     `json:"p_id,omitempty"`
	CampID    ID     `json:"c_id,omitempty"`
	Email     string `json:"email,omitempty"`
	Name      string `json:"name,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Signature string `json:"signature,omitempty"`
	Status    int    `json:"status,omitempty"`
	NickName  string `json:"nickname"`
	Title     string `json:"member_title"`
}
