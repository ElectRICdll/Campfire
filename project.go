package entity

type Project struct {
	ID          ID
	Title       string
	LeaderId    int
	Description string
	Members     map[ID]*Member
	Camps       map[ID]*Camp
	Tasks       map[ID]*Task
	FUrl        string
}

type BriefProjectDTO struct {
	ID          int    `json:"id"`
	LeaderId    int    `json:"leader"`
	Title       string `json:"title"`
	Description string `json:"des"`
	CampCount   int    `json:"camp_count"`
	TaskCount   int    `json:"task_count"`
}

type ProjectDTO struct {
	BriefProjectDTO
	Camps []CampDTO `json:"camps"`
	Tasks []Task    `json:"tasks"`
}
