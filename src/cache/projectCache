package cache

import (
	. "campfire/entity"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	ProjectCache *cache.Cache
)

func InitProjectCache() {
	ProjectCache = cache.New(30*time.Minute, 60*time.Minute)
}

func StoreProjectInCache(projID uint, project *Project) {
	ProjectCache.Set(fmt.Sprintf("%d", projID), &project, cache.DefaultExpiration)
}

func GetProjectFromCache(projID uint) (Project, bool) {
	if project, found := ProjectCache.Get(fmt.Sprintf("%d", projID)); found {
		return project.(Project), true
	} else {
		return project.(Project), false
	}

}
