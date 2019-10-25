package models

import (
	db "blog/pkg/sqlite"
	"fmt"
	"strings"
)

func (s *SuccessRecordResponse) DeleteUserRecord(id string) {
	var record RecordTable
	var check RecordTable

	db.GetDb().
		Where("id = ?", id).
		First(&record)

	if strings.EqualFold(record.CreatedAt, "") {
		fmt.Println("Record doesn't exist")
	}

	db.GetDb().
		Where("id = ?", id).
		Delete(&record)

	db.GetDb().
		Where("id = ?", id).
		First(&check)

	if strings.EqualFold(check.CreatedAt, "") {
		s.Result = "success"
	}

}
