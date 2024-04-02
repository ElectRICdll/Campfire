package entity

type Camp struct {
	ID      uint `gorm:"primaryKey;autoIncrement" json:"id"`
	ProjID  uint `gorm:"not null" json:"projID"`
	OwnerID uint `gorm:"not null" json:"ownerID"`

	Name      string `gorm:"not null" json:"name"`
	IsPrivate bool   `gorm:"not null" json:"isPrivate"`

	Owner   User     `gorm:"foreignKey:OwnerID" json:"owner"`
	Members []Member `gorm:"foreignKey:CampID;onDelete:CASCADE" json:"members"`

	Announcements  []Announcement `gorm:"foreignKey:CampID;onDelete:CASCADE" json:"announcements"`
	MessageRecords []Message      `gorm:"foreignKey:CampID;onDelete:CASCADE" json:"messages"`
}
