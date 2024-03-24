package entity

type Camp struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	ProjID  uint `gorm:"not null"`
	OwnerID uint `gorm:"not null"`

	Name      string
	IsPrivate bool
	Members   []Member `gorm:"foreignKey:CampID"`

	Announcements  []Announcement `gorm:"foreignKey:CampID"`
	MessageRecords []Message      `gorm:"foreignKey:CampID"`
}

type BriefCampDTO struct {
	ID           uint   `json:"campID" uri:"camp_id"`
	OwnerID      uint   `json:"ownerID"`
	ProjID       uint   `json:"projectID"`
	Name         string `json:"name"`
	IsPrivate    bool   `json:"isPrivate"`
	MembersCount int    `json:"memberCount"`
}

type CampDTO struct {
	ID                  uint              `json:"campID" uri:"camp_id"`
	OwnerID             uint              `json:"ownerID"`
	ProjID              uint              `json:"projectID"`
	IsPrivate           bool              `json:"isPrivate"`
	Name                string            `json:"name"`
	MembersCount        int               `json:"memberCount"`
	MembersID           []uint            `json:"membersID"`
	Members             []MemberDTO       `json:"members"`
	Announcements       []AnnouncementDTO `json:"announcements"`
	RecentMessageRecord []Message         `json:"recentMessageRecord"`
}

func (c Camp) BriefDTO() BriefCampDTO {
	return BriefCampDTO{
		ID:           c.ID,
		OwnerID:      c.OwnerID,
		ProjID:       c.ProjID,
		Name:         c.Name,
		MembersCount: len(c.Members),
	}
}

func (c Camp) DTO() CampDTO {
	return CampDTO{
		ID:            c.ID,
		OwnerID:       c.OwnerID,
		ProjID:        c.ProjID,
		Name:          c.Name,
		MembersCount:  len(c.Members),
		Members:       MembersDTO(c.Members),
		Announcements: AnnouncementsDTO(c.Announcements),
	}
}

func CampsDTO(camps []Camp) []CampDTO {
	var res []CampDTO
	for _, camp := range camps {
		res = append(res, camp.DTO())
	}

	return res
}
