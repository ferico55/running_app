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
	e.POST("/friends/requests/:id/accept", controller.AcceptFriendRequest, localMiddleware.Authorization)
	e.POST("/friends/requests/:id/decline", controller.DeclineFriendRequest, localMiddleware.Authorization)

	e.POST("/run", controller.SaveRun, localMiddleware.Authorization)
	e.GET("/run", controller.GetRunForUser, localMiddleware.Authorization)
	e.GET("/run/:id", controller.GetRunById, localMiddleware.Authorization)
	e.GET("/friends/:friendId/run/:id", controller.GetRunById, localMiddleware.Authorization)
	e.GET("/run/last", controller.GetLastRun, localMiddleware.Authorization)
	e.GET("/friends/run", controller.GetFriendsRuns, localMiddleware.Authorization)
	e.GET("/friends/:id/run", controller.GetRunForUserId, localMiddleware.Authorization)

	e.Logger.Fatal(e.Start(":80"))
}

// migrate -database "mysql://root:runner@tcp(localhost:3306)/run" -path ./migrations up
// migrate create -ext sql -dir ./migrations -seq create_users_table
