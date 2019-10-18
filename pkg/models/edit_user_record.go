package models

import (
	db "blog/pkg/sqlite"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/copier"
)

type UpdateRecordBody struct {
	Text     string `json:"text"`
	Security bool   `json:"security"`
	Subject  string `json:"subject"`
	UserId   int    `json:"user_id"`
	TagId    int    `json:"tag_id"`
	Id       int    `json:"id"`
}

func (s *SuccessRecordResponse) EditUserRecord(recordBody UpdateRecordBody) {
	var updateInfo RecordTable

	if recordBody.Security == true {
		updateInfo.Security = 1
	} else {
		updateInfo.Security = 0
	}

	copier.Copy(&updateInfo, &recordBody)

	t := time.Now().UTC().String()
	updated := strings.Split(t, " +0000 UTC")
	updateInfo.UpdatedAt = updated[0]

	var update RecordTable

	db.GetDb().
		Model(&update).
		Where("user_id = ? and id = ?", recordBody.UserId, recordBody.Id).
		Updates(map[string]interface{}{"text": updateInfo.Text, "security": updateInfo.Security, "subject": updateInfo.Subject, "updated_at": updated[0], "tag_id": updateInfo.TagId})

	var checkUpdate RecordTable
	db.GetDb().
		Raw("select * from record where user_id = $1 and id = $2", recordBody.UserId, recordBody.Id).
		Find(&checkUpdate)

	if strings.EqualFold(checkUpdate.Subject, recordBody.Subject) {
		s.Result = "success"
	} else {
		fmt.Println("update wasn't success")
	}

}
