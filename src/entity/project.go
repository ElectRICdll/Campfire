package entity

type Project struct {
	ID      ID `gorm:"primaryKey"`
	OwnerID ID

	Title       string
	Description string

	Members []*Member
	Camps   []*Camp
	Tasks   []*Task
	FUrl    string

	Owner User `gorm:"foreignKey:OwnerID"`
}

type BriefProjectDTO struct {
	ID          ID     `json:"id,omitempty"`
	OwnerID     ID     `json:"leader,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"des,omitempty"`
	CampCount   int    `json:"camp_count,omitempty"`
	TaskCount   int    `json:"task_count,omitempty"`
}

type ProjectDTO struct {
	ID          ID        `json:"id,omitempty"`
	OwnerID     ID        `json:"leader,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"des,omitempty"`
	CampCount   int       `json:"camp_count,omitempty"`
	TaskCount   int       `json:"task_count,omitempty"`
	Camps       []CampDTO `json:"camps,omitempty"`
	Tasks       []Task    `json:"tasks,omitempty"`
}
