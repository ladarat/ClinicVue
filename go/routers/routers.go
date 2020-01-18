package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	handlers "github.com/ladarat/ClinicVue/go/handles"
)

// CustomValidator type
type CustomValidator struct {
	validator *validator.Validate
}

// Validate func
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// Init initialize api routes and set up a connection.
func Init(e *echo.Echo) {

	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/login", handlers.Login())

}
