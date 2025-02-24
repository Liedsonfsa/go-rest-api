package models

import (
	"errors"
	"rest-api/db"
	"time"
)

// Event represents the data of an event
type Event struct {
	ID          int64
	Name        string		`binding:"required"`
	Description string		`binding:"required"`
	Location    string		`binding:"required"`
	DateTime    time.Time	`binding:"required"`
	UserID		int64
}

// Save saves a new event to the database
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
	e.ID = id

	return err
}

// GetAllEvents get all stored events
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

// GetEventByID gets an event with a given id
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

// Update updates a given event
func (event Event) Update() error {
	query := "UPDATE events SET event_name = ?, event_description = ?, event_location = ?, event_dateTime = ?  WHERE event_id = ?"
	
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	return err
}

// Delete remove a certain event
func (event Event) Delete() error {
	query := "DELETE FROM events WHERE event_id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}

// Register registers a user for an event
func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations (event_id, user_id) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}

// CancelRegistration remove a user's record from an event
func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}