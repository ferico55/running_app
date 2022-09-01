package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/ferico55/running_app/config"
	"github.com/ferico55/running_app/model"
	"github.com/labstack/echo/v4"
)

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization") // Bearer xxx
		authComponents := strings.Split(authHeader, " ")

		if len(authComponents) != 2 {
			return c.String(http.StatusUnauthorized, "UNAUTHORIZED")
		}

		if strings.ToLower(authComponents[0]) != "bearer" {
			return c.String(http.StatusUnauthorized, "UNAUTHORIZED")
		}

		// 2. validate & parse jwt token
		var token, err = jwt.Parse(authComponents[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(config.Secret), nil
		})
		if err != nil {
			return c.String(http.StatusUnauthorized, "UNAUTHORIZED")
		}

		//kenapa ga pake id nya aja buat dimasukin ke dalem tokennya, trus buat ngambil data lainnya masuk ke dalem db nya? soalnya ke db butuh cost gais
		if token.Valid { // udah expired belum?
			claims := token.Claims.(jwt.MapClaims)
			id := claims["id"].(string)
			name := claims["name"].(string)
			username := claims["username"].(string)
			avatarUrl := claims["avatar_url"].(string)
			idInt, _ := strconv.ParseInt(id, 10, 64)
			user := model.User{
				Id:        idInt,
				Name:      name,
				Username:  username,
				AvatarUrl: avatarUrl,
			}
			// 3. set claims (user) buat whole request
			c.Set("user", user)
			// 4. next
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "UNAUTHORIZED")
	}
}
