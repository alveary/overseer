package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// AppEngine for web engine setup
func AppEngine() *martini.ClassicMartini {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})

	return m
}

func main() {
	m := AppEngine()
	m.Run()
}
