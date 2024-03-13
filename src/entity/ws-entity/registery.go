package wsentity

type EventScope int
type EventType int

const (
	OnNobody EventScope = iota
	OnProject
	OnCamp
	OnSomeone
)

const (
	UnknownEventType EventType = iota
	TextMessageEventType
	BinaryMessageEventType
	NewTaskEventType
	NewAnnouncementEventType
	CodeGraphMessageEventType
	RequestMessageRecordEventType
	ProjectInvitationEventType
	CampInvitationEventType
)

func NewEventByType(eventType EventType) (Event, EventScope) {
	switch eventType {
	case TextMessageEventType:
		return &TextMessageEvent{}, OnCamp
	case BinaryMessageEventType:
		return &BinaryMessageEvent{}, OnCamp
	case NewTaskEventType:
		return &NewTaskEvent{}, OnProject
	case NewAnnouncementEventType:
		return &NewAnnouncementEvent{}, OnCamp
	case CodeGraphMessageEventType:
		return &CodeGraphMessageEvent{}, OnCamp
	case RequestMessageRecordEventType:
		return &RequestMessageRecordEvent{}, OnNobody
	case ProjectInvitationEventType:
		return &ProjectInvitationEvent{}, OnSomeone
	case CampInvitationEventType:
		return &CampInvitationEvent{}, OnSomeone
	default:
		return nil, OnNobody
	}
}
