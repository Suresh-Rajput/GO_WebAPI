package models

import "github.com/labstack/echo"

// StudentReq Depicts request structure to get Student details
type StudentReq struct {
	UserID int64 `query:"userID" validate:"required,numeric"`
}

// GetUsers gets user details based on userId and userName
func GetUsers(ctx echo.Context) error {
	return nil
}
