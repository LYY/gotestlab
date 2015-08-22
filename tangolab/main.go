package main

import (
	"errors"
	. "github.com/LYY/gotestlab/labcomm"
	"github.com/lunny/log"
	"github.com/lunny/tango"
	"github.com/tango-contrib/tpongo2"
	"os"
	"runtime"
)

type JsonAction struct {
	tango.Json
}

func (JsonAction) Get() interface{} {
	if true {
		return map[string]string{
			"say": "Hello tango!",
		}
	}
	return errors.New("something error")
}

type PongoAction struct {
	tpongo2.Renderer
}

func (a *PongoAction) Get() error {
	users := InitUsers()
	data := map[string]interface{}{
		"users": users,
	}
	return a.Render("pongo.html.pgo", data)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	logger := log.New(os.Stdout, "[tango] ", log.Ldefault())
	t := tango.Classic(logger)
	t.Use(tpongo2.New())
	t.Get("/json", new(JsonAction))
	t.Get("/pongo", new(PongoAction))
	t.Run(":8001")
}
