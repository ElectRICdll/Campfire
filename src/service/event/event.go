package event

type Event interface {
	Process() func() error

	ScopeID() uint
}

type EventService struct {
}
