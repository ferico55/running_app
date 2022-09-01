package controller

import (
	"database/sql"
	"encoding/json"
	"math/rand"
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
	r := rand.Intn(5)
	if r == 0 {
		requestBody.AvatarUrl = "https://cdn.dribbble.com/users/68238/screenshots/16625594/media/680d750d85b8c32b510789f523a6f906.png?compress=1&resize=400x300"
	} else if r == 1 {
		requestBody.AvatarUrl = "https://variety.com/wp-content/uploads/2021/05/Batman_Caped_Crusader2C_COLOR-e1621372528477.jpg"
	} else if r == 2 {
		requestBody.AvatarUrl = "https://flxt.tmsimg.com/assets/p22231131_i_v10_aa.jpg"
	} else if r == 3 {
		requestBody.AvatarUrl = "https://cdn.mos.cms.futurecdn.net/ceumnt7Gs3je5TVGrrYSFT.jpg"
	} else {
		requestBody.AvatarUrl = "https://www.commonsensemedia.org/sites/default/files/styles/ratio_16_9_small/public/externals/6dfe02b01e1f482eb71c74882956c859.jpg"
	}

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
