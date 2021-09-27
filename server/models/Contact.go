package models

import "gopkg.in/mgo.v2/bson"

type Contact struct {
	ID      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Number  string        `json:"number" bson:"number"`
	SavedAs string        `json:"saved_as" bson:"saved_as"`
}
