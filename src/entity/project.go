package entity

type Project struct {
	Campsites []*Campsite
	Tasks     []*Task
	CSUrl     string
}
