package entity

import "time"

type BriefCampDTO struct {
	ID     uint `json:"id" uri:"id"`
	ProjID uint `json:"projID"`

	OwnerID   uint   `json:"ownerID"`
	MembersID []uint `json:"membersID"`

	Name         string `json:"name"`
	IsPrivate    bool   `json:"isPrivate"`
	MembersCount int    `json:"membersCount"`
	Opposite     Member `json:"opposite"`
}

func (c Camp) BriefDTO() BriefCampDTO {
	return BriefCampDTO{
		ID:           c.ID,
		OwnerID:      c.OwnerID,
		ProjID:       c.ProjID,
		Name:         c.Name,
		MembersCount: len(c.Members) + 1,
	}
}

func (c Camp) BriefDTOPrivate(userID uint) BriefCampDTO {
	return BriefCampDTO{
		ID:           c.ID,
		OwnerID:      c.OwnerID,
		ProjID:       c.ProjID,
		Name:         c.Name,
		MembersCount: len(c.Members) + 1,
		Opposite: func() Member {
			if c.IsPrivate {
				if userID == c.Members[0].UserID {
					return c.Members[1]
				}
				return c.Members[0]
			}
			return Member{}
		}(),
	}
}

type BriefProjectDTO struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	BeginAt     time.Time `json:"begin"`
	Description string    `json:"description"`

	ReleaseCount int `json:"releaseCount"`
	Status       int `json:"status"`

	Branches     []string `json:"branches"`
	MembersCount int      `json:"membersCount"`
	OwnerID      uint     `json:"ownerID"`
	MembersID    []uint   `json:"membersID"`
	CampsID      []uint   `json:"campsID"`
	TasksID      []uint   `json:"tasksID"`
}

func (proj Project) BriefDTO() BriefProjectDTO {
	return BriefProjectDTO{
		ID:          proj.ID,
		Title:       proj.Title,
		Description: proj.Description,
		Status:      proj.Status,
		Branches: func(branches []Branch) []string {
			res := []string{}
			for _, branch := range branches {
				res = append(res, branch.Name)
			}
			return res
		}(proj.Branches),
		MembersCount: len(proj.Members) + 1,
	}
}

type LoginDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"username"`
	Token     string `json:"token"`
	AvatarUrl string `json:"avatarUrl"`
}
