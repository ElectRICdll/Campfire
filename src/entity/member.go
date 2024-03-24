package entity

type MemberList map[int]*Member

type Member struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	UserID uint `gorm:"not null"`
	CampID uint `gorm:"not null"`

	IsLeader bool
	Nickname string
	Title    string

	User User `gorm:"foreignKey:UserID"`
	Camp Camp `gorm:"foreignKey:CampID"`
}

type MemberDTO struct {
	ID        uint   `json:"id,omitempty" uri:"user_id" binding:"required"`
	UserID    uint   `json:"userID,omitempty"`
	ProjID    uint   `json:"projectID,omitempty"`
	CampID    uint   `json:"campID,omitempty"`
	IsLeader  bool   `json:"isLeader"`
	AvatarUrl string `json:"avatarUrl,omitempty"`
	Signature string `json:"signature,omitempty"`
	Status    int    `json:"status,omitempty"`
	Nickname  string `json:"nickname"`
	Title     string `json:"member_title"`
}

type ProjectMember struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	ProjID    uint `gorm:"not null"`
	IsCreator bool
	Title     string

	User    User    `gorm:"foreignKey:UserID"`
	Project Project `gorm:"foreignKey:ProjID"`
}

func (m Member) DTO() MemberDTO {
	return MemberDTO{
		ID:        m.ID,
		UserID:    m.UserID,
		CampID:    m.CampID,
		Nickname:  m.Nickname,
		AvatarUrl: m.User.AvatarUrl,
		Status:    m.User.Status,
		Title:     m.Title,
	}
}

func (m ProjectMember) DTO() MemberDTO {
	return MemberDTO{
		ID:        m.ID,
		UserID:    m.UserID,
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
