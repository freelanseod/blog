package addtag

import (
	"blog/pkg/models"
	"github.com/labstack/echo"
	"net/http"
)

func AddTag(c echo.Context) error {
	var addTagBody models.AddTagRequestBody

	if err := c.Bind(&addTagBody); err != nil {
		return err
	}

	addTagByUser := models.SuccessRecordResponse{}
	addTagByUser.AddTag(addTagBody)

	return c.JSON(http.StatusOK, addTagByUser)
}
