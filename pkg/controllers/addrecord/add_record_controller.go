package addrecord

import (
	"blog/pkg/models"
	"net/http"

	"github.com/labstack/echo"
)

func AddRecord(c echo.Context) error {
	var requestBody models.RecordBody

	if err := c.Bind(&requestBody); err != nil {
		return err
	}

	makeRecord := models.SuccessRecordResponse{}
	makeRecord.AddingRecord(requestBody)

	return c.JSON(http.StatusOK, makeRecord)
}
