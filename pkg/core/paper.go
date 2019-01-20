package core

import (
	"time"
)

// Paper represents a paper for discussion.
type Paper struct {
	ID       ID        `json:"id,omitempty" bson:"_id,omitempty"`
	Added    time.Time `json:"added" bson:"added"`
	URL      string    `json:"url" bson:"url"`
	Title    string    `json:"title" bson:"title"`
	Authors  []string  `json:"authors" bson:"authors"`
	Abstract string    `json:"abstract" bson:"abstract"`
}
