package models

import (
	"rest-api/db"
	"time"
)

type Event struct {
	ID          int
	Name        string		`binding:"required"`
	Description string		`binding:"required"`
	Location    string		`binding:"required"`
	DateTime    time.Time	`binding:"required"`
	UserID		int
}

var events = []Event{}

func (e *Event) Save() error {
	query := "INSERT INTO events (event_name, event_description, event_location, event_dateTime, user_id) VALUES (?, ?, ?, ?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = int(id)

	return err
}

func GetAllEvents() []Event {
	return events 
}