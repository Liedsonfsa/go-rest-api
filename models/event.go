package models

import (
	"errors"
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

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event

		if err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil 
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT event_id, event_name, event_description, event_location, event_dateTime, user_id FROM events WHERE event_id = ?"
	row, err := db.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var event Event

	if row.Next() {
		if err = row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID); err != nil {
			return nil, err
		}

		return &event, nil
	}

	return nil, errors.New("could not fetch event")
}