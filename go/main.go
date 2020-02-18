package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	mongo "github.com/ladarat/ClinicVue/go/mongo"
	repository "github.com/ladarat/ClinicVue/go/repository"
	routes "github.com/ladarat/ClinicVue/go/routers"
	"github.com/ladarat/ClinicVue/go/services"
)

func main() {

	db, err := mongo.NewDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:          middleware.DefaultSkipper,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowCredentials: true,
	}))

	patientRepo := repository.NewMongoPatientRepository(db)
	patientService := services.NewPatientService(patientRepo)

	routes := routes.Routes{
		E:       e,
		Patient: patientService,
	}

	routes.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
