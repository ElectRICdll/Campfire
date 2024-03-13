package event

import "campfire/entity"

type TextMessageEvent struct {
	entity.Message
	Content string `json:"content"`
}

func (t TextMessageEvent) Process() func() error {
	//TODO implement me
	panic("implement me")
}

func (t TextMessageEvent) ScopeID() uint {
	return t.Message.CampID
}

type BinaryMessageEvent struct {
	entity.Message
	IsChunk     bool   `json:"is_chunk,omitempty"`
	IsLastChunk bool   `json:"is_last_chunk,omitempty"`
	Binary      []byte `json:"binary"`
	Type        string `json:"type"`
}

func (b BinaryMessageEvent) Process() func() error {
	//TODO implement me
	panic("implement me")
}

func (b BinaryMessageEvent) ScopeID() uint {
	return b.Message.CampID
}

type CodeGraphMessageEvent struct {
	entity.Message
	ProjID    string `json:"p_id"`
	ObjectUrl string `json:"obj_url"`
	Lang      string `json:"lang,omitempty"`
	BeginAt   int    `json:"begin_at"`
	EndAtAt   int    `json:"end_at"`
}

func (c CodeGraphMessageEvent) Process() func() error {
	//TODO implement me
	panic("implement me")
}

func (c CodeGraphMessageEvent) ScopeID() uint {
	return c.Message.CampID
}

type FileUrlMessageEvent struct {
	entity.Message
	Url string `json:"file_url"`
}
