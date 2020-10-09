package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"mvc-go-mysql-echo/controllers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: [] string {"*"},
		AllowMethods: [] string {echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "MVC GO-MYSQL-ECHO")
	})
	e.GET("/employees", controllers.GetEmployees)
	e.Logger.Fatal(e.Start(":1234"))
}