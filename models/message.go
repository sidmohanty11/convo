package models

import "gopkg.in/mgo.v2/bson"

type Message struct {
	ID        bson.ObjectId       `json:"id" bson:"_id,omitempty"`
	To        string              `json:"to" bson:"to"`
	From      string              `json:"from" bson:"from"`
	Message   string              `json:"message" bson:"message"`
	Timestamp bson.MongoTimestamp `json:"sent_at" bson:"sent_at"`
}
