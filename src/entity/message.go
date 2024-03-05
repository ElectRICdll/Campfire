package entity

import (
	"encoding/json"
	"time"
)

const (
	UnknownMessageType = iota
	TextMessageType
	ImageMessageType
	VideoMessageType
	TaskMessageType
	AnnouncementMessageType
	CodeGraphMessageType
	DocumentMessageType
	EventMessageType
)

const (
	UnknownEvent = iota
	NewProjectEvent

	ProjectTitleChangeEvent
	MemberJoinEvent
	MemberLeftEvent
	MemberTitleChangeEvent
	NicknameChangeEvent
	CSTitleChangeEvent
)

type Message struct {
	TimeStamp  time.Time `json:"timestamp"`
	Type       int       `json:"m_type"`
	From       int       `json:"from"`
	CampsiteID int       `json:"cs"`
	TentID     int       `json:"t,omitempty"`
	ReplyID    int       `json:"r,omitempty"`
	Content    string    `json:"content"`
}

type UnknownMessage struct {
	Message
}

type TextMessage struct {
	Message
	Content string `json:"content"`
}

type BinaryMessage struct {
	Message
	Content BinaryData `json:"content"`
}

type TaskMessage struct {
	Message
	Content TaskDTO `json:"content"`
}

type EventMessage struct {
	Message
	Content struct {
		EventType int `json:"e_type"`
	} `json:"content"`
}

type AnnouncementMessage struct {
	Message
	Content AnnouncementDTO `json:"content"`
}

type CodeGraphMessage struct {
	Message
	Content struct {
		ObjectUrl string `json:"o_url"`
		Lang      string `json:"lang,omitempty"`
		Begin     int    `json:"begin"`
		End       int    `json:"end"`
	} `json:"content"`
}

type FileUrlMessage struct {
	Message
	Content string `json:"file_url"`
}

type BinaryData struct {
	IsChunk     bool   `json:"is_chunk,omitempty"`
	IsLastChunk bool   `json:"is_last_chunk,omitempty"`
	Binary      []byte `json:"binary"`
	Type        string `json:"type"`
}

func (bd *BinaryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		BinaryData
		Content interface{} `json:"content"`
	}{
		BinaryData: *bd,
		Content:    bd.Binary,
	})
}

func (bd *BinaryData) UnmarshalJSON(data []byte) error {
	aux := &struct {
		BinaryData
		Content json.RawMessage `json:"content"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	bd = &aux.BinaryData
	return json.Unmarshal(aux.Content, &bd.Binary)
}
