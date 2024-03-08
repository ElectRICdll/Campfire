package entity

type Camp struct {
	ID       ID
	Name     string
	LeaderId int
	Members  map[ID]*Member
	Anno     []*Announcement
	Belong   *Project
}

type BriefCampDTO struct {
	ID           int    `json:"id" uri:"id" binding:"required"`
	Name         string `json:"name"`
	LeaderId     int    `json:"leader"`
	MembersCount int    `json:"members_count"`
	TentsCount   int    `json:"tents_count"`
}

type CampDTO struct {
	BriefCampDTO
	Members             []MemberDTO       `json:"members"`
	Announcements       []AnnouncementDTO `json:"announcements"`
	RecentMessageRecord []Message         `json:"recent_message_record"`
}
