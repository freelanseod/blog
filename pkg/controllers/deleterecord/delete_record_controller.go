package deleterecord

import (
	"blog/pkg/models"
	"net/http"

	"github.com/labstack/echo"
)

func DeleteRecordById(c echo.Context) error {
	id := c.Param("id")

	deleteUser := models.SuccessRecordResponse{}
	deleteUser.DeleteUserRecord(id)
	
	return c.JSON(http.StatusOK, deleteUser)
}
