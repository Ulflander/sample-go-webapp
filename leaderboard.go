package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"time"
)

const (
	resultPerPage = 20
)

func Leaderboard() string {
	return "Welcome to the leaderboard"
}

func GetParamPage(params martini.Params) (int, int) {
	if params["page"] == "" {
		return 1, 0
	}

	page, err := strconv.Atoi(params["page"])
	if err != nil {
		return 1, 0
	}

	if page < 1 {
		page = 1
	}

	return page, resultPerPage * (page - 1)
}

func FindLeader(name string, db *mgo.Database) Leader {
	result := Leader{}
	db.C("leaders").Find(bson.M{"n": name}).One(&result)
	return result
}

func GetLeader(params martini.Params, res http.ResponseWriter, r render.Render, db *mgo.Database) {
	leader := FindLeader(params["name"], db)
	if leader.ID == "" {
		r.JSON(404, nil)
	} else {
		// Find rank
		rank, err := db.C("leaders").Find(bson.M{"s": bson.M{"$gt": leader.Score}}).Count()
		if err == nil {
			leader.Rank = rank + 1
		}
		r.JSON(200, leader)
	}
}

func GetLeaders(params martini.Params, res http.ResponseWriter, r render.Render, db *mgo.Database) {
	var leaders []Leader
	skip := 0

	_, skip = GetParamPage(params)

	err := db.C("leaders").Find(nil).Sort("-s", "-ts").Skip(skip).Limit(resultPerPage).All(&leaders)

	if err != nil {
		r.JSON(500, err)
		return
	}

	i := skip + 1
	for index := range leaders {
		leaders[index].Rank = i
		i += 1
	}
	r.JSON(200, leaders)
}

func PostLeader(leader Leader, res http.ResponseWriter, r render.Render, db *mgo.Database) {
	// Check for existing leader
	existing := FindLeader(leader.Name, db)

	// Not exists, create it
	if existing.ID == "" {
		leader.ID = bson.NewObjectId()
		leader.Timestamp = time.Now()
		err := db.C("leaders").Insert(leader)
		if err == nil {
			r.JSON(200, leader)
		} else {
			r.JSON(500, nil)
		}

		// Exists, update it
	} else {
		existing.Score = leader.Score
		existing.Timestamp = time.Now()
		err := db.C("leaders").Update(bson.M{"_id": bson.ObjectId(existing.ID)}, existing)

		if err == nil {
			r.JSON(200, existing)
		} else {
			r.JSON(500, nil)
		}
	}

}
