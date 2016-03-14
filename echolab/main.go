package main

import (
	. "github.com/LYY/gotestlab/labcomm"
	"github.com/echo-contrib/pongor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	// "github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"runtime"
)

// Handler
func PongoGetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		users := InitUsers()
		data := map[string]interface{}{
			"users": users,
		}
		return c.Render(200, "pongo.html.pgo", data)
	}
}

func StringHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(200, "hello girl")
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	serv := echo.New()
	serv.Use(middleware.Logger())
	serv.Use(middleware.Recover())
	r := pongor.GetRenderer()
	serv.SetRenderer(r)
	serv.Get("/pongo", PongoGetHandler())
	serv.Get("/string", StringHandler())

	serv.Run(standard.New(":8002"))
	// serv.Run(fasthttp.New(":8002"))
}
