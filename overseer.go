package main

import (
	"log"
	"net/http"

	"github.com/alveary/overseer/registry"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

type requestedService struct {
	name    string
	address string
}

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

	m.Get("/alive", func(r render.Render) {
		r.Status(200)
	})

	m.Post("/", binding.Form(requestedService{}), func(errors binding.Errors, req requestedService, resp http.ResponseWriter) {
		s, err := services.Register(req.name, req.address)

		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Print(s)
	})

	return m
}

func init() {
	log.Print("Initializing Overseer Instance:")
}

func main() {
	m := AppEngine()
	m.Run()
}
