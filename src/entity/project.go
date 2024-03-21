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
	ID      uint `gorm:"primaryKey;autoIncrement"`
	OwnerID uint `gorm:"not null"`

	Title        string `gorm:"not null"`
	Description  string
	Status       int
	BeginAt      time.Time
	Branches     []Branch      `gorm:"foreignKey:ProjID"`
	PullRequests []PullRequest `gorm:"foreignKey:ProjID"`
	Releases     []Release     `gorm:"foreignKey:ProjID"`

	Members []ProjectMember `gorm:"foreignKey:ProjID"`
	Camps   []Camp          `gorm:"foreignKey:ProjID"`
	Tasks   []Task          `gorm:"foreignKey:ProjID"`
	FUrl    string
}

type BriefProjectDTO struct {
	ID           uint      `json:"projectID,omitempty"`
	OwnerID      uint      `json:"ownerID,omitempty"`
	Title        string    `json:"projectTitle,omitempty"`
	BeginAt      time.Time `json:"begin"`
	Description  string    `json:"description,omitempty"`
	BranchCount  int       `json:"branchCount"`
	PullCount    int       `json:"pullCount"`
	ReleaseCount int       `json:"releaseCount"`
	CampCount    int       `json:"campCount,omitempty"`
	TaskCount    int       `json:"taskCount,omitempty"`
	Status       int       `json:"status"`
}

type ProjectDTO struct {
	ID           uint          `json:"projectID,omitempty"`
	OwnerID      uint          `json:"ownerID,omitempty"`
	Title        string        `json:"projectTitle,omitempty"`
	Description  string        `json:"description,omitempty"`
	BeginAt      time.Time     `json:"begin"`
	Branches     []Branch      `json:"branches"`
	PullRequests []PullRequest `json:"pull_requests"`
	Releases     []Release     `json:"releases"`
	Camps        []CampDTO     `json:"camps,omitempty"`
	Tasks        []TaskDTO     `json:"tasks,omitempty"`
	Status       int           `json:"status"`
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
		ID:           p.ID,
		OwnerID:      p.OwnerID,
		Title:        p.Title,
		Description:  p.Description,
		Branches:     p.Branches,
		PullRequests: p.PullRequests,
		Releases:     p.Releases,
		Camps:        CampsDTO(p.Camps),
		Tasks:        TasksDTO(p.Tasks),
	}
}

func ProjectsDTO(projects []Project) []ProjectDTO {
	res := []ProjectDTO{}
	for _, project := range projects {
		res = append(res, project.DTO())
	}
	return res
}

type Branch struct {
	ProjID uint
	Name   string `json:"branch"`
}

const (
	Open = iota
	Closed
	Merged
)

type PullRequest struct {
	ID      string `json:"id" gorm:"primaryKey"`
	ProjID  uint
	Title   string `json:"title"`
	Body    string `json:"body"`
	Branch  string `json:"branch"`
	OwnerID uint   `json:"ownerID"`
	Status  int
}

type Release struct {
	ProjID   uint
	Version  string    `json:"version"`
	Date     time.Time `json:"date"`
	FilePath string    `json:"-"`
}
