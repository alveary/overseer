package main

import (
	"flag"
	"strconv"

	"github.com/go-martini/martini"
)

// AppEngine for web engine setup
func AppEngine() *martini.ClassicMartini {
	m := martini.Classic()

	return m
}

func main() {
	var port int
	flag.IntVar(&port, "p", 9001, "the port number")
	flag.Parse()

	m := AppEngine()
	m.RunOnAddr(":" + strconv.Itoa(port))
}
