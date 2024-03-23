package cache

import (
	. "campfire/entity"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	CampCache *cache.Cache
)

func InitCampCache() {
	CampCache = cache.New(30*time.Minute, 60*time.Minute)
}

func StoreCampInCache(campID uint, camp *Camp) {
	CampCache.Set(fmt.Sprintf("%d", campID), &camp, cache.DefaultExpiration)
}

func GetCampFromCache(campID uint) (Camp, bool) {
	if camp, found := CampCache.Get(fmt.Sprintf("%d", campID)); found {
		return camp.(Camp), true
	} else {
		return camp.(Camp), false
	}

}
