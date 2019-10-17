package models

import (
	db "blog/pkg/sqlite"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/copier"
)

type RecordBody struct {
	Text     string `json:"text"`
	Security bool   `json:"security"`
	Subject  string `json:"subject"`
	UserId   int    `json:"user_id"`
	TagId    int    `json:"tag_id"`
}

type SuccessRecordResponse struct {
	Result string `json:"result"`
}

type RecordTable struct {
	Id        int
	Text      string
	Security  int
	Subject   string
	UserId    int
	CreatedAt string
	UpdatedAt string
	TagId     int
}

func (RecordTable) TableName() string {
	return "record"
}

func (s *SuccessRecordResponse) AddingRecord(recordBody RecordBody) {
	var insertInfo RecordTable
	if recordBody.Security == true {
		insertInfo.Security = 1
	} else {
		insertInfo.Security = 0
	}

	copier.Copy(&insertInfo, &recordBody)

	t := time.Now().UTC().String()
	created := strings.Split(t, " +0000 UTC")
	updated := strings.Split(t, " +0000 UTC")
	insertInfo.CreatedAt = created[0]
	insertInfo.UpdatedAt = updated[0]

	db.GetDb().
		Create(&insertInfo)

	var checkInsert RecordTable

	db.GetDb().
		Raw("select * from record where subject = $1 and user_id = $2", recordBody.Subject, recordBody.UserId).
		Find(&checkInsert)

	if strings.EqualFold(checkInsert.CreatedAt, created[0]) {
		s.Result = "success"
	} else {
		fmt.Println("insert wasn't successful")
	}

}
