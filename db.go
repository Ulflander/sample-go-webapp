package main

import (
	"github.com/go-martini/martini"
	"gopkg.in/mgo.v2"
)

var MongoDB mgo.Database

func InitializeMongo() martini.Handler {
	sess, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	s := sess.Clone()
	defer s.Close()
	EnsureDBIndexes(s.DB("go_test"))

	return func(c martini.Context) {
		s := sess.Clone()
		c.Map(s.DB("go_test"))
		defer s.Close()
		c.Next()
	}
}
