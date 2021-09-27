package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Message struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username  string        `json:"username" bson:"username"`
	Content   string        `json:"content" bson:"content"`
	Timestamp time.Time     `json:"timestamp" bson:"timestamp"`
}
