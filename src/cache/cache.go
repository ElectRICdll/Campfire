package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var ProjectCacheInstance *BaseCache
var UserCacheInstance *BaseCache

//type ProjectCacheWrapper struct {
//
//}
//
//func StoreProject(project Project) {
//
//}

type BaseCache struct {
	*cache.Cache
	IsInitialized bool
}

func InitCacheInstance() {
	ProjectCacheInstance.Cache = cache.New(24*time.Hour, 24*7*time.Hour)
	ProjectCacheInstance.IsInitialized = true

	UserCacheInstance.Cache = cache.New(24*7*time.Hour, 24*30*time.Hour)
	UserCacheInstance.IsInitialized = true
}
