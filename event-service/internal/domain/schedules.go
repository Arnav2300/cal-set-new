package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Schedule struct {
	ID          bson.ObjectID `json:"_id" bson:"_id"`
	EventId     bson.ObjectID `json:"event_id" bson:"event_id"`
	StartTime   time.Time     `json:"start_time" bson:"start_time"`
	EndTime     time.Time     `json:"end_time" bson:"end_time"`
	IsAvailable string        `json:"is_available" bson:"is_available"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}
