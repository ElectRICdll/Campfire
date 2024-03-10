package entity

import "time"

const (
	OnUnknown = iota
	OnProject
	OnCamp
	OnSomeone
)

var EventTypeIndex = []struct {
	InnerType Event
	Scope     int
}{
	{TextMessageEvent{}, OnCamp},
	{BinaryMessageEvent{}, OnCamp},
	{BinaryMessageEvent{}, OnCamp},
	{TaskEvent{}, OnProject},
	{AnnouncementEvent{}, OnCamp},
	{CodeGraphMessageEvent{}, OnCamp},
	{BinaryMessageEvent{}, OnCamp},
}

type Event interface {
	Execute() func() error
	ScopeID() uint
}

type TextMessageEvent struct {
	Message
	Content string `json:"content"`
}

func (t TextMessageEvent) Execute() func() error {
	//TODO implement me
	panic("implement me")
}

func (t TextMessageEvent) ScopeID() uint {
	return t.Message.CampID
}

type BinaryMessageEvent struct {
	Message
	IsChunk     bool   `json:"is_chunk,omitempty"`
	IsLastChunk bool   `json:"is_last_chunk,omitempty"`
	Binary      []byte `json:"binary"`
	Type        string `json:"type"`
}

func (b BinaryMessageEvent) Execute() func() error {
	//TODO implement me
	panic("implement me")
}

func (b BinaryMessageEvent) ScopeID() uint {
	return b.Message.CampID
}

type CodeGraphMessageEvent struct {
	Message
	ProjID    string `json:"p_id"`
	ObjectUrl string `json:"obj_url"`
	Lang      string `json:"lang,omitempty"`
	BeginAt   int    `json:"begin_at"`
	EndAtAt   int    `json:"end_at"`
}

func (c CodeGraphMessageEvent) Execute() func() error {
	//TODO implement me
	panic("implement me")
}

func (c CodeGraphMessageEvent) ScopeID() uint {
	return c.Message.CampID
}

type TaskEvent struct {
	Timestamp time.Time `json:"timestamp"`
	TaskDTO
}

func (t TaskEvent) Execute() func() error {
	//TODO implement me
	panic("implement me")
}

func (t TaskEvent) ScopeID() uint {
	return t.TaskDTO.ProjID
}

type AnnouncementEvent struct {
	Timestamp time.Time `json:"timestamp"`
	AnnouncementDTO
}

func (a AnnouncementEvent) Execute() func() error {
	//TODO implement me
	panic("implement me")
}

func (a AnnouncementEvent) ScopeID() uint {
	return a.AnnouncementDTO.CampID
}

type FileUrlMessage struct {
	Message
	Url string `json:"file_url"`
}
