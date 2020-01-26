package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	handlers "github.com/ladarat/ClinicVue/go/handles"
	"github.com/ladarat/ClinicVue/go/services"
)

// CustomValidator type
type CustomValidator struct {
	validator *validator.Validate
}

type Routes struct {
	E       *echo.Echo
	Patient services.PatientService
}

// Validate func
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// Init initialize api routes and set up a connection.

func (r *Routes) Init() {

	r.E.Validator = &CustomValidator{validator: validator.New()}

	r.E.POST("/login", handlers.Login())
	r.E.POST("/patient", handlers.CreatePatient(r.Patient))

}
