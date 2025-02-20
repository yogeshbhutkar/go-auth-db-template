package models

import (
	"time"

	"github.com/yogeshbhutkar/go-jwt-with-db-template/db"
)

type Event struct {
	ID          int64
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	insertEvent := `
		INSERT INTO events(name, description, location, dateTime, user_id)
		VALUES  (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(insertEvent)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.DateTime, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetEvents() []Event {
	return events
}
