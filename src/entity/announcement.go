package entity

import "time"

type Announcement struct {
	ID      ID
	OwnerID int
	CampID  int
	Title   string
	Begin   time.Time
	Content string
}

type AnnouncementDTO struct {
	ID      ID        `json:"ID"`
	CampID  int       `json:"c_id"`
	OwnerID int       `json:"o_id"`
	Title   string    `json:"title"`
	Begin   time.Time `json:"begin"`
	Content string    `json:"content"`
}

func AnnouncementsToDTO(annos []Announcement) []*AnnouncementDTO {
	var res []*AnnouncementDTO

	for _, anno := range annos {
		dto := &AnnouncementDTO{
			ID:      anno.ID,
			OwnerID: anno.OwnerID,
			Title:   anno.Title,
			Begin:   anno.Begin,
			Content: anno.Content,
		}
		res = append(res, dto)
	}

	return res
}
