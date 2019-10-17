package addrecord

import (
	"blog/pkg/models"
	"net/http"

	"github.com/labstack/echo"
)

func AddRecord(c echo.Context) (err error) {
	var requestBody models.RecordBody

	if error := c.Bind(&requestBody); error != nil {
		return error
	}

	makeRecord := models.SuccessRecordResponse{}
	makeRecord.AddingRecord(requestBody)

	return c.JSON(http.StatusOK, makeRecord)
}
