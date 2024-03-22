package ws

const (
	OnNobody int = iota
	OnProject
	OnCamp
	OnSomeone
)

const (
	UnknownEventType int = iota
	AuthEventType
	PingEventType
	TextMessageEventType
	BinaryMessageEventType
	CodeGraphMessageEventType
	RequestMessageRecordEventType

	ProjectInfoChangedEventType
	NewTaskEventType
	NewAnnouncementEventType
	CampInfoChangedEventType
	CampDisableEventType
	MemberInfoChangedEventType
	MemberExitedEventType

	ProjectInvitationEventType
	CampInvitationEventType
)

var EventsByType = map[int]func() (Event, int){
	UnknownEventType:              func() (Event, int) { return nil, 0 },
	TextMessageEventType:          func() (Event, int) { return &TextMessageEvent{}, OnCamp },
	BinaryMessageEventType:        func() (Event, int) { return &BinaryMessageEvent{}, OnCamp },
	CodeGraphMessageEventType:     func() (Event, int) { return &CodeGraphMessageEvent{}, OnCamp },
	RequestMessageRecordEventType: func() (Event, int) { return &RequestMessageRecordEvent{}, OnNobody },

	ProjectInfoChangedEventType: func() (Event, int) { return &ProjectInfoChangedEvent{}, OnProject },
	NewTaskEventType:            func() (Event, int) { return &NewTaskEvent{}, OnProject },
	NewAnnouncementEventType:    func() (Event, int) { return &NewAnnouncementEvent{}, OnCamp },
	CampInfoChangedEventType:    func() (Event, int) { return &CampInfoChangedEvent{}, OnCamp },
	CampDisableEventType:        func() (Event, int) { return &CampDisableEvent{}, OnCamp },
	MemberInfoChangedEventType:  func() (Event, int) { return &MemberInfoChangedEvent{}, OnCamp },
	MemberExitedEventType:       func() (Event, int) { return &MemberExitedEvent{}, OnCamp },

	ProjectInvitationEventType: func() (Event, int) { return &ProjectInvitationEvent{}, OnSomeone },
	CampInvitationEventType:    func() (Event, int) { return &CampInvitationEvent{}, OnSomeone },
}

func ScopeByType(eventType int) int {
	_, res := EventsByType[eventType]()
	return res
}
