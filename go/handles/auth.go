package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// LoginRequest struct
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse struct
type LoginResponse struct {
	UserID string `json:"userID"`
}

// Login by POST /login
func Login() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		u := LoginRequest{}
		if err = c.Bind(&u); err != nil {
			return err
		}
		if err = c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		ur := LoginResponse{UserID: "aa23"}

		return c.JSON(http.StatusOK, ur)

	}
}
