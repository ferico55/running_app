package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ferico55/running_app/model"
	"github.com/ferico55/running_app/repository"
	"github.com/labstack/echo/v4"
)

func SaveRun(c echo.Context) error {
	user := c.Get("user").(model.User)
	decoder := json.NewDecoder(c.Request().Body)
	var requestBody model.Run
	err := decoder.Decode(&requestBody)
	if err != nil {
		return responseError(c, http.StatusUnprocessableEntity, err.Error())
	}

	err = repository.SaveRun(requestBody, user.Id, c.Request().Context())
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	return responseJson(c, http.StatusCreated, nil)
}

func GetRunForUser(c echo.Context) error {
	user := c.Get("user").(model.User)

	runs, err := repository.GetRunForUser(user.Id, c.Request().Context())
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	return responseJson(c, http.StatusOK, runs)
}

func GetRunById(c echo.Context) error {
	stringId := c.Param("id")
	runId, err := strconv.ParseInt(stringId, 10, 64)
	if err != nil {
		return responseError(c, http.StatusUnprocessableEntity, err.Error())
	}

	run, err := repository.GetRunById(runId, c.Request().Context())
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	return responseJson(c, http.StatusOK, run)
}

func GetLastRun(c echo.Context) error {
	user := c.Get("user").(model.User)

	run, err := repository.GetLastRunForUser(user.Id, c.Request().Context())
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	return responseJson(c, http.StatusOK, run)
}

func GetFriendsRuns(c echo.Context) error {
	user := c.Get("user").(model.User)
	friends, err := repository.GetFriendForUser(user.Id, c.Request().Context())
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	runs, err := repository.GetRunFromUsers(friends, c.Request().Context())
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	return responseJson(c, http.StatusOK, runs)
}

func GetRunForUserId(c echo.Context) error {
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

	runs, err := repository.GetRunForUser(friendId, c.Request().Context())
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	return responseJson(c, http.StatusOK, runs)
}
