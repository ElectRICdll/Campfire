package event

const (
	OnUnknown = iota
	OnNobody
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
	{RequestMessageRecordEvent{}, OnNobody},
	{ProjectInvitationEvent{}, OnSomeone},
	{CampInvitationEvent{}, OnSomeone},
}
