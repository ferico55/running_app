package controller

import (
	"net/http"

	"github.com/ferico55/running_app/model"
	"github.com/ferico55/running_app/repository"
	"github.com/labstack/echo/v4"
)

func GetFriends(c echo.Context) error {
	user := c.Get("user").(model.User)
	friends, err := repository.GetFriendForUser(user.Id, c.Request().Context())
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	return responseJson(c, http.StatusOK, friends)
}
