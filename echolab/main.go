package main

import (
	. "github.com/LYY/gotestlab/labcomm"
	"github.com/echo-contrib/pongor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"runtime"
)

func PongoGet(ctx *echo.Context) error {
	users := InitUsers()
	data := map[string]interface{}{
		"users": users,
	}
	return ctx.Render(200, "pongo.html.pgo", data)
}

func StringGet(ctx *echo.Context) error {
	return ctx.String(200, "%s", "hello girl")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	serv := echo.New()
	serv.Use(middleware.Logger())
	serv.Use(middleware.Recover())
	r := pongor.GetRenderer()
	serv.SetRenderer(r)
	serv.Get("/pongo", PongoGet)
	serv.Get("/string", StringGet)

	serv.Run(":8002")
}
