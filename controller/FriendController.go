package controller

import (
	"database/sql"
	"net/http"
	"strconv"

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

func AddFriend(c echo.Context) error {
	user := c.Get("user").(model.User)
	stringId := c.Param("id")
	friendId, err := strconv.ParseInt(stringId, 10, 64)
	if err != nil {
		return responseError(c, http.StatusUnprocessableEntity, err.Error())
	}

	err = repository.IsFriended(user.Id, friendId, c.Request().Context())
	if err == nil {
		return responseError(c, http.StatusUnprocessableEntity, "This user is your friend already or you still have pending invite.")
	} else if err != sql.ErrNoRows {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	err = repository.AddFriend(user.Id, friendId, c.Request().Context())
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	return responseJson(c, http.StatusCreated, nil)
}

func RemoveFriend(c echo.Context) error {
	user := c.Get("user").(model.User)
	stringId := c.Param("id")
	friendId, err := strconv.ParseInt(stringId, 10, 64)
	if err != nil {
		return responseError(c, http.StatusUnprocessableEntity, err.Error())
	}

	err = repository.IsFriended(user.Id, friendId, c.Request().Context())
	if err == sql.ErrNoRows {
		return responseError(c, http.StatusUnprocessableEntity, "You are not friend of this user")
	} else if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	err = repository.RemoveFriend(user.Id, friendId, c.Request().Context())
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	return responseJson(c, http.StatusOK, nil)
}

func GetPendingFriendRequest(c echo.Context) error {
	user := c.Get("user").(model.User)
	friends, err := repository.GetPendingFriendRequest(user.Id, c.Request().Context())
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	return responseJson(c, http.StatusOK, friends)
}

func AcceptFriendRequest(c echo.Context) error {
	user := c.Get("user").(model.User)
	stringId := c.Param("id")
	friendId, err := strconv.ParseInt(stringId, 10, 64)
	if err != nil {
		return responseError(c, http.StatusUnprocessableEntity, err.Error())
	}

	err = repository.IsFriended(user.Id, friendId, c.Request().Context())
	if err == sql.ErrNoRows {
		return responseError(c, http.StatusUnprocessableEntity, "You have no request from this user")
	} else if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	err = repository.AcceptFriendRequest(user.Id, friendId, c.Request().Context())
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	return responseJson(c, http.StatusOK, nil)
}

func DeclineFriendRequest(c echo.Context) error {
	user := c.Get("user").(model.User)
	stringId := c.Param("id")
	friendId, err := strconv.ParseInt(stringId, 10, 64)
	if err != nil {
		return responseError(c, http.StatusUnprocessableEntity, err.Error())
	}

	err = repository.IsFriended(user.Id, friendId, c.Request().Context())
	if err == sql.ErrNoRows {
		return responseError(c, http.StatusUnprocessableEntity, "You have no request from this user")
	} else if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	err = repository.RemoveFriend(user.Id, friendId, c.Request().Context())
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	return responseJson(c, http.StatusOK, nil)
}
