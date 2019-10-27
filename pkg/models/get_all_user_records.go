package models

import (
	db "blog/pkg/sqlite"
	"fmt"
)

type Record struct {
	Id        int    `json:"id"`
	Text      string `json:"text"`
	Security  bool   `json:"security"`
	Subject   string `json:"subject"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	TagId     int    `json:"tag_id"`
}

type AllRecordsResponse struct {
	Result string   `json:"result"`
	Data   []Record `json:"data"`
}

type CountRows struct {
	Query int `gorm:"column:count(*)"`
}

func (a *AllRecordsResponse) GetAllUserRecords(UserId string) {
	var countR CountRows

	db.GetDb().
		Raw("select count(*) from record where user_id = $1", UserId).
		Find(&countR)

	a.Data = make([]Record, countR.Query)

	rows, err := db.GetDb().
		Model(&Record{}).
		Raw("select * from record where user_id = $1", UserId).
		Rows()
	defer rows.Close()

	if err != nil {
		fmt.Println("Couldn't get records err=%+v", err)
	}

	count := 0
	a.Result = "success"

	for rows.Next() {
		db.ScanDbRows(rows, &a.Data[count])
		count++
	}

}
