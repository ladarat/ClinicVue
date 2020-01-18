package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ladarat/ClinicVue/go/model"
)

// Login func
func Login() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		u := model.LoginRequest{}
		if err = c.Bind(&u); err != nil {
			return err
		}
		if err = c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		ur := model.LoginResponse{UserID: "aa23"}

		return c.JSON(http.StatusOK, ur)

	}
}
