package domain

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type EventStatus string

const (
	StatusActive      EventStatus = "Active"
	StatusFillingFast EventStatus = "Filling Fast"
	StatusClosed      EventStatus = "Closed"
)

type Event struct {
	ID          bson.ObjectID `json:"id" bson:"_id"`
	OrganizerId string        `json:"organizer_id" bson:"organizer_id"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"desc" bson:"desc"`
	Location    string        `json:"loc" bson:"loc"`
	Seats       int           `json:"seats" bson:"seats"`
	Status      string        `json:"status" bson:"status"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
	Tags        []string      `json:"tags" bson:"tags"`
}

func (e *Event) ValidateStatus() error {
	switch e.Status {
	case string(StatusActive), string(StatusFillingFast), string(StatusClosed):
		return nil
	default:
		return errors.New("invalid status value")
	}
}
