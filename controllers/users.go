package controllers

import (
	"GO_WebAPI/models"
	"net/http"

	"github.com/labstack/echo"
)

// GetUsers gets user details based on userId and userName
func GetUsers(ctx echo.Context) error {

	req := new(models.StudentReq)
	err := ctx.Bind(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request")
	}
	err = ctx.Validate(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request")
	}
	return ctx.JSON(http.StatusOK, "Successfully fetched student user details")
}
