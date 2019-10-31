package getuserrecord

import (
	"blog/pkg/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetRecordById(c echo.Context) error {
	id := c.Param("id")
	recordId := c.Param("recordid")

	recordUser := models.GetRecordByRecordId(id, recordId)

	return c.JSON(http.StatusOK, recordUser)
}
