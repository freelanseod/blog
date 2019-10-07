package auth

import (
	"blog/pkg/models"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

func Authentication(c echo.Context) (err error) {
	var authRequestBody models.AuthRequestBody

	if error := c.Bind(&authRequestBody); error != nil {
		return error
	}

	if strings.EqualFold(authRequestBody.Phone, "") && (authRequestBody.Email != "") && (authRequestBody.Password != "") {
		auth := models.AuthenticationResponse{}
		auth.AuthWithEmail(authRequestBody)

		return c.JSON(http.StatusOK, auth)
	} else if strings.EqualFold(authRequestBody.Password, "") {
		auth := models.ErrorResponse{}
		auth.EmptyPassword(authRequestBody)

		return c.JSON(http.StatusUnauthorized, auth)
	} else if strings.EqualFold(authRequestBody.Email, "") && (authRequestBody.Phone != "") && (authRequestBody.Password != "") {
		auth := models.AuthenticationResponse{}
		auth.AuthWithPhone(authRequestBody)

		return c.JSON(http.StatusOK, auth)
	} else if strings.EqualFold(authRequestBody.Email, "") || strings.EqualFold(authRequestBody.Phone, "") {
		auth := models.ErrorResponse{}
		auth.EmptyLogin(authRequestBody)

		return c.JSON(http.StatusUnauthorized, auth)
	}

	return err
}
