package game

type EventType int32

const (
	Won  EventType = 0
	Lost EventType = 1
)

type Event struct {
	UserId    uint64
	EventType EventType
}
