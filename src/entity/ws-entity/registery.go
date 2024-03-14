package wsentity

type EventScope int
type EventType int

const (
	OnNobody int = iota
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
	CampInfoChangedEventType
	CampDisableEventType
	MemberInfoChangedEventType
	MemberExitedEventType
	CodeGraphMessageEventType
	RequestMessageRecordEventType
	ProjectInvitationEventType
	CampInvitationEventType
)

func GetEventByType(eventType EventType) (Event, int) {
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
	case CampDisableEventType:
		return &CampDisableEvent{}, OnCamp
	case MemberInfoChangedEventType:
		return &MemberInfoChangedEvent{}, OnCamp
	default:
		return nil, OnNobody
	}
}
