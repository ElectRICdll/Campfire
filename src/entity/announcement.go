package entity

import "time"

type Announcement struct {
	ID      uint      `gorm:"primaryKey;autoIncrement"`
	OwnerID uint      `gorm:"not null"`
	ProjID  uint      `gorm:"not null"`
	CampID  uint      `gorm:"not null"`
	Title   string    `gorm:"not null"`
	Begin   time.Time `gorm:"not null"`
	Content string    `gorm:"not null"`
}

type AnnouncementDTO struct {
	ID      uint      `json:"ID"`
	CampID  uint      `json:"c_id"`
	ProjID  uint      `json:"p_id"`
	OwnerID uint      `json:"o_id"`
	Title   string    `json:"title"`
	Begin   time.Time `json:"begin"`
	Content string    `json:"content"`
}

func (a Announcement) DTO() AnnouncementDTO {
	return AnnouncementDTO{
		ID:      a.ID,
		OwnerID: a.OwnerID,
		ProjID:  a.ProjID,
		CampID:  a.CampID,
		Title:   a.Title,
		Begin:   a.Begin,
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
