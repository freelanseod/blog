package editrecord

import (
	"blog/pkg/models"
	"net/http"

	"github.com/labstack/echo"
)

func EditRecord(c echo.Context) (err error) {
	var requestBody models.UpdateRecordBody

	if error := c.Bind(&requestBody); error != nil {
		return error
	}

	changeRecord := models.SuccessRecordResponse{}
	changeRecord.EditUserRecord(requestBody)

	return c.JSON(http.StatusOK, changeRecord)
}
