package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alveary/overseer/registry"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

// ServiceRegistry provides access to the dummy Registry
func ServiceRegistry() registry.Registry {
	registry := registry.Registry{make(map[string][]registry.Service)}
	return registry
}

// AppEngine for web engine setup
func AppEngine() *martini.ClassicMartini {
	m := martini.Classic()
	servicereg := ServiceRegistry()
	m.Get("/", binding.Form(registry.Service{}), func(errors binding.Errors, service registry.Service, resp http.ResponseWriter) {
		fmt.Println(errors)
		servicereg.Register(service)
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
