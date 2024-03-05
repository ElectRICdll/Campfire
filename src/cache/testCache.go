package cache

import "campfire/entity"

var TestProjects = []entity.Project{
	{
		Campsites: []*entity.Campsite{
			{
				ID:   1,
				Name: "老登交流群",
				Members: map[entity.ID]*entity.Member{
					TestUsers[1].ID: {
						TestUsers[1],
						"刘新宇",
						"后端开发",
					},
					TestUsers[2].ID: {
						TestUsers[2],
						"姚佳铭",
						"前端开发",
					},
					TestUsers[3].ID: {
						TestUsers[3],
						"江梓豪",
						"",
					},
				},
			},
		},
		Tasks: []*entity.Task{
			{
				ID:         1,
				CreatorID:  1,
				ReceiverID: []entity.ID{2, 3},
				Content:    "完成",
				Status:     1,
			},
			{
				ID:         2,
				CreatorID:  1,
				ReceiverID: []entity.ID{2, 3},
				Content:    "未完成",
				Status:     1,
			},
		},
		CSUrl: "",
	},
}

var TestUsers = map[entity.ID]*entity.User{
	1: {
		ID:       1,
		Email:    "hare@email.com",
		Name:     "electric",
		IsOnline: false,
	},
	2: {
		ID:       2,
		Email:    "bdeng@email.com",
		Name:     "wryte",
		IsOnline: false,
	},
	3: {
		ID:       2,
		Email:    "koishi@email.com",
		Name:     "koishi",
		IsOnline: false,
	},
}
