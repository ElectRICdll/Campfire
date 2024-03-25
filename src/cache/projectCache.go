package cache

import (
	. "campfire/entity"
	"errors"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	ProjectCache *cache.Cache
)

func Instance() {

}

func InitProjectCache() {
	ProjectCache = cache.New(30*time.Minute, 60*time.Minute)
}

func StoreProjectInCache(projID uint, project *Project) {
	ProjectCache.Set(fmt.Sprintf("%d", projID), &project, cache.DefaultExpiration)
}

func GetProjectFromCache(projID uint) (*Project, bool) {
	if project, found := ProjectCache.Get(fmt.Sprintf("%d", projID)); found {
		return project.(*Project), true
	} else {
		return project.(*Project), false
	}

}

func StoreTaskToProject(projID uint, task Task) error {
	if project, found := ProjectCache.Get(fmt.Sprintf("%d", projID)); found {
		project.(*Project).Tasks = append(project.(*Project).Tasks, task)
	}
	return errors.New("no such data in cache")
}

func GetTaskFromCache(projID, taskID uint) (*Task, error) {
	if project, found := ProjectCache.Get(fmt.Sprintf("%d", projID)); found {
		for _, task := range project.(*Project).Tasks {
			if task.ID == taskID {
				return &task, nil
			}
		}
		return nil, errors.New("no such data in cache")
	}
	return nil, errors.New("no such data in cache")
}

func EditTaskFromCache(update Task) error {
	if project, found := ProjectCache.Get(fmt.Sprintf("%d", update.ProjID)); found {
		for _, task := range project.(*Project).Tasks {
			t := &task
			if t.ID == update.ID {
				if update.Title != "" {
					t.Title = update.Title
				}
				if !update.BeginAt.IsZero() {
					t.BeginAt = update.BeginAt
					task.Stop()
					task.StartATimer()
				}
				if !update.EndAt.IsZero() {
					t.EndAt = update.EndAt
					task.Stop()
					task.StartATimer()
				}
				if update.Content != "" {
					t.Content = update.Content
				}
				if update.Status != 0 {
					t.Status = update.Status
				}
			}
		}
	}
	return errors.New("no such data in cache")
}

func DelTaskFromCache(del Task) error {
	if project, found := ProjectCache.Get(fmt.Sprintf("%d", del.ProjID)); found {
		for index, task := range project.(*Project).Tasks {
			t := &task
			if t.ID == del.ID {
				project.(*Project).Tasks = append(project.(*Project).Tasks[0:index], project.(*Project).Tasks[index+1:]...)
				return nil
			}
		}
		return errors.New("no such data in cache")
	}
	return errors.New("no such data in cache")
}
