package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alveary/overseer/registry"
	"github.com/alveary/overseer/watchdog"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

// ServiceRegistry provides access to the dummy Registry
func ServiceRegistry() registry.Registry {
	registry := registry.Registry{make(map[string]registry.Service)}
	return registry
}

// AppEngine for web engine setup
func AppEngine() *martini.ClassicMartini {
	m := martini.Classic()
	m.Use(render.Renderer())

	servicereg := ServiceRegistry()

	m.Post("/", binding.Json(registry.Service{}), func(errors binding.Errors, service registry.Service, resp http.ResponseWriter) {
		fmt.Println(errors)
		servicereg.Register(service)
		watchdog.Watch(&service)
	})

	m.Get("/service/:name", func(r render.Render, params martini.Params) {
		r.JSON(200, servicereg.Services[params["name"]])

	})

	return m
}

func main() {
	var port int
	flag.IntVar(&port, "p", 8999, "the port number")
	flag.Parse()

	m := AppEngine()
	m.RunOnAddr(":" + strconv.Itoa(port))
}
