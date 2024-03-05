package entity

import "sync"

type Campsite struct {
	ID      ID
	Name    string
	Members map[ID]*Member
	Ams     []*Announcement

	Tents map[ID]*Tent

	Lock sync.Mutex
}
