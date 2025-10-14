package models

import "time"

type Events struct {
	Id          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserId      int
}

var events = []Events{}

func (e Events) Save() {
	//Add it to database
	events = append(events, e)
}

func GetAllEvents() []Events {
	return events
}
