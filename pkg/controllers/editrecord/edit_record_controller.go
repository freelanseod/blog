package editrecord

import (
	"blog/pkg/models"
	"net/http"

	"github.com/labstack/echo"
)

func EditRecord(c echo.Context) error {
	var requestBody models.UpdateRecordBody

	if err := c.Bind(&requestBody); err != nil {
		return err
	}

	changeRecord := models.SuccessRecordResponse{}
	changeRecord.EditUserRecord(requestBody)

	return c.JSON(http.StatusOK, changeRecord)
}
