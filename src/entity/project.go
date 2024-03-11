package entity

type Project struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	OwnerID uint `gorm:"not null"`

	Title       string `gorm:"not null"`
	Description string

	Members []ProjectMember `gorm:"foreignKey:ProjID"`
	Camps   []Camp          `gorm:"many2many:project_camps;"`
	Tasks   []Task          `gorm:"foreignKey:ProjID"`
	FUrl    string
}

type BriefProjectDTO struct {
	ID          uint   `json:"p_id,omitempty"`
	OwnerID     uint   `json:"o_id,omitempty"`
	Title       string `json:"project_title,omitempty"`
	Description string `json:"des,omitempty"`
	CampCount   int    `json:"camp_count,omitempty"`
	TaskCount   int    `json:"task_count,omitempty"`
}

type ProjectDTO struct {
	ID          uint      `json:"p_id,omitempty"`
	OwnerID     uint      `json:"o_id,omitempty"`
	Title       string    `json:"project_title,omitempty"`
	Description string    `json:"des,omitempty"`
	CampCount   int       `json:"camp_count,omitempty"`
	TaskCount   int       `json:"task_count,omitempty"`
	Camps       []CampDTO `json:"camps,omitempty"`
	Tasks       []TaskDTO `json:"tasks,omitempty"`
}

func (p Project) BriefDTO() BriefProjectDTO {
	return BriefProjectDTO{
		ID:          p.ID,
		OwnerID:     p.OwnerID,
		Title:       p.Title,
		Description: p.Description,
		CampCount:   len(p.Camps),
		TaskCount:   len(p.Tasks),
	}
}

func (p Project) DTO() ProjectDTO {
	return ProjectDTO{
		ID:          p.ID,
		OwnerID:     p.OwnerID,
		Title:       p.Title,
		Description: p.Description,
		Camps:       CampsDTO(p.Camps),
		CampCount:   len(p.Camps),
		Tasks:       TasksDTO(p.Tasks),
		TaskCount:   len(p.Tasks),
	}
}

func ProjectsDTO(projects []Project) []ProjectDTO {
	res := []ProjectDTO{}
	for _, project := range projects {
		res = append(res, project.DTO())
	}
	return res
}
