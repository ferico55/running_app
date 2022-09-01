package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/ferico55/running_app/model"
	"github.com/ferico55/running_app/repository"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	var requestBody model.RegisterRequest
	err := decoder.Decode(&requestBody)

	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	err = repository.IsUsernameFound(requestBody.Username, c.Request().Context())
	if err == nil {
		return responseError(c, http.StatusNotFound, "Sorry, the username is unavailable :(")
	} else if err != sql.ErrNoRows {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), 12)
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}
	requestBody.Password = string(hashedPassword[:])
	requestBody.AvatarUrl = "https://cdn.dribbble.com/users/68238/screenshots/16625594/media/680d750d85b8c32b510789f523a6f906.png?compress=1&resize=400x300"

	userID, err := repository.RegisterUser(requestBody, c.Request().Context())
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}

	userGotten, err := repository.GetUserById(userID, c.Request().Context())
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}
	return responseJson(c, http.StatusOK, userGotten)
}
