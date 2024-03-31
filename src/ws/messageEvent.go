package ws

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
	IsChunk     bool   `json:"isChunk,omitempty"`
	IsLastChunk bool   `json:"isLastChunk,omitempty"`
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
	ProjID    string `json:"projID"`
	ObjectUrl string `json:"objectUrl"`
	Lang      string `json:"lang,omitempty"`
	BeginAt   int    `json:"begin"`
	EndAtAt   int    `json:"end"`
}

func (c CodeGraphMessageEvent) ToMessage() entity.Message {
	return c.Message
}

func (c CodeGraphMessageEvent) ScopeID() uint {
	return c.Message.CampID
}

type MarkdownMessageEvent struct {
	entity.Message
	Content string `json:"content"`
}

func (c MarkdownMessageEvent) ToMessage() entity.Message {
	return c.Message
}

func (c MarkdownMessageEvent) ScopeID() uint {
	return c.Message.CampID
}

type FileUrlMessageEvent struct {
	entity.Message
	Url string `json:"fileUrl"`
}
