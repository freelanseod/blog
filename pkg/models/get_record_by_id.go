package models

import (
	db "blog/pkg/sqlite"
)

type RecordResponse struct {
	Result string `json:"result"`
	Record Record `json:"record"`
}

func GetRecordByRecordId(id string, recordId string) interface{} {
	var r RecordResponse
	db.GetDb().
		Raw("select * from record where user_id = $1 and id = $2", id, recordId).
		Find(&r.Record)

	switch r.Record.Text {
	case "":
		var e ErrorResponse
		e.Result = "unsuccess"
		e.Data.ErrorInformation = "Record doesn't exist"
		e.Data.ErrorMessage = "unknown_record_id"
		return e
	default:
		r.Result = "success"
		return r
	}

}
