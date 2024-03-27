package entity

import "time"

const (
	Planning = iota
	Developing
	Testing
	Releasing
	Released
	Maintaining
	Archived
)

type Project struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"projectID"`

	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Status      int       `gorm:"not null" json:"status"`
	BeginAt     time.Time `json:"begin"`

	Branches []Branch  `gorm:"foreignKey:ProjID;onDelete:CASCADE" json:"branches"`
	Releases []Release `gorm:"foreignKey:ProjID;onDelete:CASCADE" json:"releases"`

	OwnerID uint `json:"-"`
	Owner   User `gorm:"foreignKey:OwnerID;onDelete:CASCADE" json:"owner"`

	Members []ProjectMember `gorm:"foreignKey:ProjID;onDelete:CASCADE" json:"members"`

	Camps []Camp `gorm:"foreignKey:ProjID;onDelete:CASCADE" json:"camps"`
	Tasks []Task `gorm:"foreignKey:ProjID;onDelete:CASCADE" json:"tasks"`
}

const (
	Open = iota
	Closed
	Merged
)

type Branch struct {
	ID          uint     `gorm:"primaryKey;autoIncrement" json:"-"`
	ProjID      uint     `gorm:"not null" json:"projID"`
	OwnerID     uint     `gorm:"not null" json:"ownerID"`
	Name        string   `gorm:"not null;unique" json:"branch"`
	Commits     []Commit `gorm:"foreignKey:ProjID;onDelete:CASCADE" json:"-"`
	CommitCount int      `gorm:"-" json:"commitCount"`
	Status      int      `gorm:"not null" json:"status"`
}

type Commit struct {
	ID          string `gorm:"primaryKey;autoIncrement" json:"-"`
	ProjID      uint   `gorm:"not null" json:"projID"`
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	OwnerID     uint   `json:"ownerID"`
}

type Release struct {
	ID       string    `gorm:"primaryKey;autoIncrement" json:"id"`
	ProjID   uint      `gorm:"not null" json:"projID"`
	Version  string    `gorm:"not null;unique" json:"version"`
	Date     time.Time `gorm:"not null" json:"date"`
	FilePath string    `gorm:"not null;uri" json:"-"`
}
