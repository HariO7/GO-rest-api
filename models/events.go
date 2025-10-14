package models

import (
	"time"

	db "example.com/rest-api/database"
)

type Event struct {
	Id          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	Date        time.Time
	UserId      int
}

func (e Event) Save() error {
	query := `INSERT INTO events(name, description, location, date, user_id)
	VALUES(?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.Date, e.UserId)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()

	e.Id = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
