package events

type EventType int

const (
	GENERAL EventType = iota
	NOTICE
	SUCCESS
	WARNING
	ERROR
)

type Event struct {
	Type    EventType
	Code    int
	Message string
}
