package server

import (
	"GO_WebAPI/controllers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Initialize used to initialize http server
func Initialize() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	//ROUTES
	e.GET("/status", checkStatus)
	g := e.Group("/users") // creating users routing group
	//g.Use() // Authenticate  request

	g.POST("/info", controllers.GetUsers)
	// Server
	e.Logger.Fatal(e.Start(":8080"))
}
func checkStatus(echo echo.Context) error {
	return echo.HTML(http.StatusOK, "<br> Hello World !! </br>")
}
