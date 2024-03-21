package entity

import "time"

type Announcement struct {
	ID      uint      `gorm:"primaryKey;autoIncrement"`
	OwnerID uint      `gorm:"not null"`
	ProjID  uint      `gorm:"not null"`
	CampID  uint      `gorm:"not null"`
	Title   string    `gorm:"not null"`
	BeginAt time.Time `gorm:"not null"`
	Content string    `gorm:"not null"`
}

type AnnouncementDTO struct {
	ID      uint      `json:"id"`
	CampID  uint      `json:"campID"`
	ProjID  uint      `json:"projectID"`
	OwnerID uint      `json:"ownerID"`
	Title   string    `json:"title"`
	BeginAt time.Time `json:"begin"`
	Content string    `json:"content"`
}

func (a Announcement) DTO() AnnouncementDTO {
	return AnnouncementDTO{
		ID:      a.ID,
		OwnerID: a.OwnerID,
		ProjID:  a.ProjID,
		CampID:  a.CampID,
		Title:   a.Title,
		BeginAt: a.BeginAt,
		Content: a.Content,
	}
}

func AnnouncementsDTO(anno []Announcement) []AnnouncementDTO {
	var res []AnnouncementDTO

	for _, anno := range anno {
		res = append(res, anno.DTO())
	}

	return res
}
