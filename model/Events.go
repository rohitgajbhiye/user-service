package model

type Event struct {
	Event string `json:"event"`
	Scope string `json:"scope"`
}

type EventsWrapper struct {
	Events []Event `json:"events"`
}
