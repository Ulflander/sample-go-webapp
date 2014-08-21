package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Charset: "UTF-8",
	}))

	m.Use(func(res http.ResponseWriter) {
		res.Header().Set("X-Built-With", "Go")
	})

	m.Use(InitializeMongo())

	SetupRoutes(m)

	m.RunOnAddr("localhost:8020")
}
