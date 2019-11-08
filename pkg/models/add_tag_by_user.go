package models

import (
	db "blog/pkg/sqlite"
	"strings"
)

type AddTagRequestBody struct {
	Name string `json:"name"`
}

type TagsTable struct {
	Id int 
	Name string
}

func (TagsTable) TableName() string {
	return "tags"
}

func (s *SuccessRecordResponse) AddTag(addTagBody AddTagRequestBody) {
	var name AddTagRequestBody
	db.GetDb().
	Raw("select name from tags where name = $1", addTagBody.Name).
	Find(&name)

	if strings.EqualFold(name.Name, "") {
		var tags TagsTable
		tags.Name = addTagBody.Name
		
		db.GetDb().
		Create(&tags)

		s.Result = "success"
	}

}
