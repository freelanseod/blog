package getrecords

import (
	"blog/pkg/models"
	"github.com/labstack/echo"
	"net/http"
)

func GetRecordsByUser(c echo.Context) error {
	UserId := c.Param("id")

	recordsList := models.AllRecordsResponse{}
	recordsList.GetAllUserRecords(UserId)

	return c.JSON(http.StatusOK, recordsList)
}
