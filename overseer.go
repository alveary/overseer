package main

import (
	"log"
	"net/http"

	"github.com/alveary/overseer/registry"
	"github.com/alveary/overseer/service"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

// ServiceRegistry provides access to the dummy Registry
func ServiceRegistry() registry.Registry {
	registry := registry.Registry{make(map[string]interface{})}
	registry.UnleashWatchdogs()
	return registry
}

// AppEngine for web engine setup
func AppEngine() *martini.ClassicMartini {
	m := martini.Classic()
	m.Use(render.Renderer())

	servicereg := ServiceRegistry()

	m.Post("/", binding.Json(service.Service{}), func(errors binding.Errors, service service.Service, resp http.ResponseWriter, log *log.Logger) {
		log.Printf("registering new Service: %s", service.Name)
		servicereg.Register(&service)
	})

	m.Get("/", func(r render.Render) {
		r.JSON(200, servicereg.Services)

	})

	m.Get("/:name", func(r render.Render, params martini.Params) {
		registered := servicereg.Services[params["name"]]

		if registered.(*service.Service).Root != "" {
			r.JSON(200, registered)
		} else {
			r.JSON(404, service.Service{})
		}
	})

	return m
}

func main() {
	m := AppEngine()
	m.Run()
}
