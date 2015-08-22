package main

import (
	. "github.com/LYY/gotestlab/labcomm"
	"github.com/echo-contrib/pongor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	serv := echo.New()
	serv.Use(middleware.Logger())
	serv.Use(middleware.Recover())
	r := pongor.GetRenderer()
	serv.SetRenderer(r)
	serv.Get("/pongo", func(ctx *echo.Context) error {
		users := InitUsers()
		data := map[string]interface{}{
			"users": users,
		}
		ctx.Render(200, "pongo.html.pgo", data)
		return nil
	})

	serv.Run(":8002")
}
