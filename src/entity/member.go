package entity

type MemberList map[int]*Member

type Member struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	ProjID uint `gorm:"index;not null"`
	UserID uint `gorm:"index;not null"`
	CampID uint `gorm:"index;not null"`

	IsLeader bool
	Nickname string
	Title    string

	User User `gorm:"foreignKey:UserID"`
	Camp Camp `gorm:"foreignKey:CampID"`
}

type MemberDTO struct {
	ID        uint   `json:"id,omitempty" uri:"user_id" binding:"required"`
	UserID    uint   `json:"u_id,omitempty"`
	ProjID    uint   `json:"p_id,omitempty"`
	CampID    uint   `json:"c_id,omitempty"`
	Email     string `json:"email,omitempty"`
	Name      string `json:"name,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Signature string `json:"signature,omitempty"`
	Status    int    `json:"status,omitempty"`
	NickName  string `json:"nickname"`
	Title     string `json:"member_title"`
}
