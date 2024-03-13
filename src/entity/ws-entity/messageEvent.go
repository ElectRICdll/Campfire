package wsentity

import (
	"campfire/entity"
)

type MessageEvent interface {
	ToMessage() entity.Message
}

type TextMessageEvent struct {
	entity.Message
	Content string `json:"content"`
}

func (t TextMessageEvent) ToMessage() entity.Message {
	return t.Message
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

func (b BinaryMessageEvent) ToMessage() entity.Message {
	return b.Message
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

func (c CodeGraphMessageEvent) ToMessage() entity.Message {
	return c.Message
}

func (c CodeGraphMessageEvent) ScopeID() uint {
	return c.Message.CampID
}

type FileUrlMessageEvent struct {
	entity.Message
	Url string `json:"file_url"`
}
