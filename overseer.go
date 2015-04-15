package main

import (
	"log"

	"github.com/alveary/overseer/registry"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// AppEngine for web engine setup
func AppEngine() *martini.ClassicMartini {
	m := martini.Classic()
	m.Use(render.Renderer())

	services, err := registry.New()

	if err != nil {
		log.Printf("RegistryConnectionFailed: %s", err)
	}

	m.Get("/", func(r render.Render) {
		all := services.All()
		r.HTML(200, "index", all)
	})

	return m
}

func init() {
	log.Print("Initializing Overseer Instance")
}

func main() {
	m := AppEngine()
	m.Run()
}
