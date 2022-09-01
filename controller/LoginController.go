package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ferico55/running_app/config"
	"github.com/ferico55/running_app/model"
	"github.com/ferico55/running_app/repository"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponseStruct struct {
	User         model.User `json:"user,omitEmpty"`
	Token        string     `json:"token,omitEmpty"`
	RefreshToken string     `json:"refresh_token,omitEmpty"`
}

func Login(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	var requestBody loginRequest
	// decode vs encode
	err := decoder.Decode(&requestBody)
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
		// 500 status code
	}

	user, err := repository.GetUserByUsername(requestBody.Username, c.Request().Context())

	if user == nil {
		return responseJson(c, http.StatusUnprocessableEntity, "Invalid username")
	}
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err.Error())
	}
	fmt.Println(user.Password)
	fmt.Println(requestBody.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password))
	if err != nil {
		// fmt.Println(err)
		return responseJson(c, http.StatusUnprocessableEntity, "Wrong password woi!")
	}

	accessToken, refreshToken := GenerateAuthToken(*user)
	response := loginResponseStruct{
		User:         *user,
		Token:        accessToken,
		RefreshToken: refreshToken,
	}

	return responseJson(c, http.StatusOK, response)
}

func GenerateAuthToken(user model.User) (string, string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         strconv.FormatInt(user.Id, 10),
		"name":       user.Name,
		"username":   user.Username,
		"avatar_url": user.AvatarUrl,
		"exp":        time.Now().Add(time.Hour * 24 * 30).Unix(), // sebulan
	})
	tokenString, _ := token.SignedString([]byte(config.Secret))

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  strconv.FormatInt(user.Id, 10),
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
	})
	refreshTokenString, _ := refreshToken.SignedString([]byte(config.Secret))

	return tokenString, refreshTokenString
}
