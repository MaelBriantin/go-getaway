package game

type EventType struct {
	Name        string
	Description string
}

type Event struct {
	Name        string
	Description string
	Type				EventType
}