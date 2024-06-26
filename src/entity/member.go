package entity

import "time"

type MemberList map[int]*Member

type Member struct {
	UserID uint `gorm:"primaryKey;autoIncrement:false" json:"userID"`
	CampID uint `gorm:"primaryKey;autoIncrement:false" json:"campID"`

	Nickname string    `json:"nickname"`
	Title    string    `json:"title"`
	IsRuler  bool      `gorm:"not null" json:"isRuler"`
	LastRead time.Time `gorm:"" json:"lastRead"`

	User User `gorm:"foreignKey:UserID;onDelete:CASCADE" json:"user"`
}

type ProjectMember struct {
	UserID uint `gorm:"primaryKey" json:"userID"`
	ProjID uint `gorm:"primaryKey" json:"projID"`

	Title string `json:"title"`

	ReceivingTasks []TaskReceivers `gorm:"foreignKey:MemberUserID;references:UserID;onDelete:SET NULL"`
	ExecutingTasks []TaskExecutors `gorm:"foreignKey:MemberUserID;references:UserID;onDelete:SET NULL"`

	User User `gorm:"foreignKey:UserID;onDelete:CASCADE" json:"user"`
}
