package main

import (
	"github.com/ferico55/running_app/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	localMiddleware "github.com/ferico55/running_app/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("login", controller.Login)
	e.POST("register", controller.RegisterUser)

	e.GET("/friends", controller.GetFriends, localMiddleware.Authorization)
	e.POST("/friends/:id/add", controller.AddFriend, localMiddleware.Authorization)
	e.DELETE("/friends/:id/remove", controller.RemoveFriend, localMiddleware.Authorization)
	e.GET("/friends/requests", controller.GetPendingFriendRequest, localMiddleware.Authorization)

	e.Logger.Fatal(e.Start(":8888"))
}

// migrate -database "mysql://root:runner@tcp(localhost:3306)/run" -path ./migrations up
// migrate create -ext sql -dir ../migrations -seq create_users_table
