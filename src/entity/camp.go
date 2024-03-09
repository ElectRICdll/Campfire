package entity

type Camp struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	ProjID  uint `gorm:"not null"`
	OwnerID uint `gorm:"not null"`

	Name    string
	Members []Member `gorm:"many2many:member_camps"`

	Announcements  []Announcement `gorm:"foreignKey:CampID"`
	MessageRecords []Message      `gorm:"foreignKey:CampID"`
}

type BriefCampDTO struct {
	ID           uint   `json:"id" uri:"id" binding:"required"`
	OwnerID      uint   `json:"leader"`
	ProjID       uint   `json:"p_id"`
	Name         string `json:"name"`
	MembersCount int    `json:"members_count"`
}

type CampDTO struct {
	ID                  uint              `json:"id" uri:"id" binding:"required"`
	OwnerID             uint              `json:"leader"`
	ProjID              uint              `json:"p_id"`
	Name                string            `json:"name"`
	MembersCount        int               `json:"members_count"`
	Members             []MemberDTO       `json:"members"`
	Announcements       []AnnouncementDTO `json:"announcements"`
	RecentMessageRecord []Message         `json:"recent_message_record"`
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
