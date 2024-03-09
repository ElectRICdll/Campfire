package entity

type Project struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	OwnerID uint `gorm:"not null"`

	Title       string `gorm:"not null"`
	Description string

	Members []Member `gorm:"foreignKey:ProjID"`
	Camps   []Camp   `gorm:"many2many:project_camps;"`
	Tasks   []Task   `gorm:"foreignKey:ProjID"`
	FUrl    string
}

type BriefProjectDTO struct {
	ID          uint   `json:"id,omitempty"`
	OwnerID     uint   `json:"leader,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"des,omitempty"`
	CampCount   int    `json:"camp_count,omitempty"`
	TaskCount   int    `json:"task_count,omitempty"`
}

type ProjectDTO struct {
	ID          uint      `json:"id,omitempty"`
	OwnerID     uint      `json:"leader,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"des,omitempty"`
	CampCount   int       `json:"camp_count,omitempty"`
	TaskCount   int       `json:"task_count,omitempty"`
	Camps       []CampDTO `json:"camps,omitempty"`
	Tasks       []Task    `json:"tasks,omitempty"`
}
