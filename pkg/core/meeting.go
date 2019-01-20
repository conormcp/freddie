package core

import (
	"time"
)

// Meeting represents a meeting.
type Meeting struct {
	ID       ID        `json:"id,omitempty" bson:"_id,omitempty"`
	Date     time.Time `json:"date" bson:"date"`
	Speaker  string    `json:"speaker" bson:"speaker"`
	Title    string    `json:"title" bson:"title"`
	Affil    string    `json:"affil" bson:"affil"`
	Abstract string    `json:"abstract,omitempty" bson:"abstract,omitempty"`
}
