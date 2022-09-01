package controller

import (
	"github.com/labstack/echo/v4"
)

type baseResponse struct {
	Success bool        `json:"is_success"`
	Error   *string     `json:"error_message"`
	Data    interface{} `json:"data"`
}

func responseJson(c echo.Context, status int, data interface{}) error {
	response := baseResponse{true, nil, data}
	return c.JSON(status, response)
}

func responseError(c echo.Context, status int, errorMessage string) error {
	response := baseResponse{false, &errorMessage, nil}
	return c.JSON(status, response)
}
