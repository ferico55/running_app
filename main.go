package main

import (
	"github.com/ferico55/running_app/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("login", controller.Login)
	e.POST("register", controller.RegisterUser)

	e.Logger.Fatal(e.Start(":8888"))
}

// migrate -database "mysql://root:runner@tcp(localhost:3306)/run" -path ./migrations up
