package entity

import "time"

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
	ID      ID `json:"m_id" gorm:"primaryKey"`
	OwnerID ID `json:"from"`
	ProjID  ID `json:"p_id"`
	CampID  ID `json:"c_id"`
	ReplyID ID `json:"r_id"`

	Timestamp time.Time `json:"timestamp"`
	Type      int       `json:"m_type"`
	Content   string    `json:"content"`

	Owner User    `json:"-" gorm:"foreignKey:FromID"`
	Proj  Project `json:"-" gorm:"foreignKey:ProjID"`
	Camp  Camp    `json:"-" gorm:"foreignKey:CampID"`
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
