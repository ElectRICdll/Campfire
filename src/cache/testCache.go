package cache

import (
	. "campfire/entity"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	userCache *cache.Cache
)

func InitCache() {
	userCache = cache.New(30*time.Minute, 60*time.Minute) // 设置缓存，过期时间为 30 分钟，清理间隔为 60 分钟
}

func StoreUserInCache(user User) {
	userCache.Set(fmt.Sprintf("user:%d", user.ID), user, cache.DefaultExpiration)
}

func GetUserFromCache(userID int) (bool, User) {
	if user, found := userCache.Get(fmt.Sprintf("user:%d", userID)); found {
		return true, user.(User)
	}
	return false, User{}
}

var (
	TaskCacheByUserID *cache.Cache
)

func TaskInitByUserID() {
	TaskCacheByUserID = cache.New(30*time.Minute, 60*time.Minute) // 设置缓存，过期时间为 30 分钟，清理间隔为 60 分钟
}

func StoreTaskInCacheByUserID(user User, task Task) {
	TaskCacheByUserID.Set(fmt.Sprintf("user:%d", user.ID), task, cache.DefaultExpiration)
}

func GetTaskFromCacheByUserID(userID int) (bool, Task) {
	if task, found := userCache.Get(fmt.Sprintf("user:%d", userID)); found {
		return true, task.(Task)
	}
	return false, Task{}
}

var (
	messageCache *cache.Cache
)

func MsgInit() {
	messageCache = cache.New(30*time.Minute, 60*time.Minute) // 设置缓存，过期时间为 30 分钟，清理间隔为 60 分钟
}

func StoreMessageInCache(userID int, message Message) {
	messages, found := messageCache.Get(fmt.Sprintf("user:%d", userID))
	if !found {
		messages = make([]Message, 0)
	}
	messages = append(messages.([]Message), message)
	messageCache.Set(fmt.Sprintf("user:%d", userID), messages, cache.DefaultExpiration)
}

func GetMessagesFromCache(userID int) []Message {
	if messages, found := messageCache.Get(fmt.Sprintf("user:%d", userID)); found {
		return messages.([]Message)
	}
	return nil
}

var (
	taskCache *cache.Cache
)

func TaskInit() {
	taskCache = cache.New(30*time.Minute, 60*time.Minute) // 设置缓存，过期时间为 30 分钟，清理间隔为 60 分钟
}

//func StoreTaskInCache(user []User, task Task) {
//	taskCache.Set(fmt.Sprintf("task:%d", task.ID), user, cache.DefaultExpiration)
//}
//
//func GetUnfinishedUsersFromCache(taskID int) []User {
//	if users, found := taskCache.Get(fmt.Sprintf("task:%d", taskID)); found {
//		return users.([]User)
//	}
//	return nil
//}

//var TestProjects = []entity.Project{
//	{
//		Camps: map[uint]*entity.Camp{
//			1: {
//				ID:   1,
//				Name: "老登交流群",
//				Members: map[uint]*entity.Member{
//					TestUsers[1].ID: {
//						TestUsers[1],
//						"刘新宇",
//						"后端开发",
//					},
//					TestUsers[2].ID: {
//						TestUsers[2],
//						"姚佳铭",
//						"前端开发",
//					},
//					TestUsers[3].ID: {
//						TestUsers[3],
//						"江梓豪",
//						"",
//					},
//				},
//			},
//		},
//		FUrl: "",
//	},
//}
