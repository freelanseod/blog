package registrationuser

import (
	"blog/pkg/models"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func Registration(c echo.Context) error {
	var user models.RegistrationRequestBody

	if err := c.Bind(&user); err != nil {
		return err
	}

	lenght := len(user.Phone)

	if !(strings.ContainsAny(user.Email, "@") && lenght == 12) {
		var e models.ErrorResponse

		e.Result = "unsuccess"
		e.Data.ErrorInformation = "Check parameters for registration"
		e.Data.ErrorMessage = "invalid_parametr"

		return c.JSON(http.StatusBadRequest, e)
	} 
		registration := models.UserRegistrationResponse{}
		registration.RegistrationUser(user)

		return c.JSON(http.StatusOK, registration)
}
