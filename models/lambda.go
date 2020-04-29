package models

import "gopkg.in/mgo.v2/bson"

type Lambda struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	UserID int
	Tag string
	Generic interface{}
}
