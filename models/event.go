package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	// TODO: Add it to the DB.
	events = append(events, e)
}

func GetEvents() []Event {
	return events
}
