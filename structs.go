package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	Leader struct {
		ID        bson.ObjectId `json:"id"     bson:"_id,omitempty"`
		Name      string        `json:"name"   bson:"n"               binding:"required"`
		Score     int           `json:"score"  bson:"s"               binding:"required"`
		Rank      int           `json:"rank"`
		Timestamp time.Time     `json:"ts"     bson:"ts"`
	}
)

func EnsureDBIndexes(db *mgo.Database) {
	db.C("leaders").EnsureIndex(mgo.Index{
		Key:        []string{"n", "s"},
		Unique:     true,
		DropDups:   true,
		Background: true, // See notes.
		Sparse:     true,
	})
}
