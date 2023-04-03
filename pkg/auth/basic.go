package auth

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
	"pluseid.io/invitation/config"
)

func BasicAuthentication(username, password string, c echo.Context) (bool, error) {
	if subtle.ConstantTimeCompare([]byte(username), []byte(config.GetAdmin().Username)) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte(config.GetAdmin().Password)) == 1 {
		return true, nil
	}
	return false, nil
}
