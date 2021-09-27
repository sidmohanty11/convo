package models

import "gopkg.in/mgo.v2/bson"

type Chat struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name         string        `json:"name" bson:"name"`
	Participants []string      `json:"participants" bson:"participants"`
	Messages     []*Message    `json:"messages" bson:"messages"`
}
