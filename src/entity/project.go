package entity

import (
	"time"
)

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

	Branches []Branch `gorm:"foreignKey:ProjID;onDelete:CASCADE" json:"branches"`

	OwnerID uint `json:"-"`
	Owner   User `gorm:"foreignKey:OwnerID;onDelete:CASCADE" json:"owner"`

	Members []ProjectMember `gorm:"foreignKey:ProjID;onDelete:CASCADE" json:"members"`

	Camps []Camp `gorm:"foreignKey:ProjID;onDelete:CASCADE" json:"camps"`
	Tasks []Task `gorm:"foreignKey:ProjID;onDelete:CASCADE" json:"tasks"`

	Path string `json:"-"`
}

const (
	Open = iota
	Closed
	Merged
)

type Branch struct {
	Name   string `gorm:"not null;" json:"branch"`
	ProjID uint   `gorm:"not null" json:"projID"`
}

type Release struct {
	ID       string    `gorm:"primaryKey;autoIncrement" json:"id"`
	ProjID   uint      `gorm:"not null" json:"projID"`
	Version  string    `gorm:"not null" json:"version"`
	Date     time.Time `gorm:"not null" json:"date"`
	FilePath string    `gorm:"not null" json:"-"`
}
