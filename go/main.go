package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	routes "github.com/ladarat/ClinicVue/go/routers"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:          middleware.DefaultSkipper,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowCredentials: true,
	}))

	routes.Init(e)

	e.Logger.Fatal(e.Start(":1323"))
}
