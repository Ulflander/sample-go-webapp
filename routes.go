package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

func SetupRoutes(m *martini.ClassicMartini) {
	m.Get("/", Leaderboard)
	m.Get("/leaders", GetLeaders)
	m.Get("/leaders/:page", GetLeaders)
	m.Get("/leader/:name", GetLeader)
	m.Post("/leader", binding.Json(Leader{}), binding.ErrorHandler, PostLeader)
}
