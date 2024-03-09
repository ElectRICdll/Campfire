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
	Name      string `json:"name,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Signature string `json:"signature,omitempty"`
	Status    int    `json:"status,omitempty"`
	NickName  string `json:"nickname"`
	Title     string `json:"member_title"`
}

type ProjectMember struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"index;not null"`
	ProjID    uint `gorm:"index;not null"`
	IsCreator bool
	Title     string

	User    User    `gorm:"foreignKey:UserID"`
	Project Project `gorm:"foreignKey:ProjID"`
}

func (m Member) DTO() MemberDTO {
	return MemberDTO{
		ID:        m.ID,
		UserID:    m.UserID,
		ProjID:    m.ProjID,
		CampID:    m.CampID,
		NickName:  m.Nickname,
		AvatarUrl: m.User.AvatarUrl,
		Status:    m.User.Status,
		Title:     m.Title,
	}
}

func (m ProjectMember) DTO() MemberDTO {
	return MemberDTO{
		ID:        m.ID,
		UserID:    m.UserID,
		ProjID:    m.ProjID,
		AvatarUrl: m.User.AvatarUrl,
		Status:    m.User.Status,
		Title:     m.Title,
	}
}

func MembersDTO(members []Member) []MemberDTO {
	res := []MemberDTO{}
	for _, member := range members {
		res = append(res, member.DTO())
	}

	return res
}
