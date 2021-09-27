package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username"`
	Number   string        `json:"number" bson:"number"`
	Password string        `json:"-" bson:"password"`
	LastSeen time.Time     `json:"last_seen" bson:"last_seen"`
	Token    string        `json:"token,omitempty" bson:"-"`
	Contacts []*Contact    `json:"contacts" bson:"contacts"`
	Chats    []*Chat       `json:"chats" bson:"chats"`
}
