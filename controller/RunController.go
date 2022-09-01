package controller

import (
	"encoding/json"
	"net/http"

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
