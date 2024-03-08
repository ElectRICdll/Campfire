package entity

type Camp struct {
	ID      ID `gorm:"primaryKey"`
	ProjID  ID
	OwnerID ID

	Name    string
	Members []Member
	Anno    []Announcement

	Proj  Project `gorm:"foreignKey:ProjID"`
	Owner User    `gorm:"foreignKey:OwnerID"`
}

type BriefCampDTO struct {
	ID           ID     `json:"id" uri:"id" binding:"required"`
	OwnerID      ID     `json:"leader"`
	Name         string `json:"name"`
	MembersCount int    `json:"members_count"`
}

type CampDTO struct {
	ID                  ID                `json:"id" uri:"id" binding:"required"`
	OwnerID             ID                `json:"leader"`
	Name                string            `json:"name"`
	MembersCount        int               `json:"members_count"`
	Members             []MemberDTO       `json:"members"`
	Announcements       []AnnouncementDTO `json:"announcements"`
	RecentMessageRecord []Message         `json:"recent_message_record"`
}
