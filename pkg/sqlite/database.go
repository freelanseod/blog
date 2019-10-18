package sqlite

import (
	"blog/pkg/sqlite/seeds"
	"blog/pkg/sqlite/tables"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db, _ = gorm.Open("sqlite3", "./database/database.sqlite")

func Initialize() {
	runMigrations()
}

func GetDb() *gorm.DB {
	return db
}

func runMigrations() {
	db := GetDb()

	db.Exec(tables.CreateUserTable)
	db.Exec(tables.CreateTagsTable)
	db.Exec(tables.CreateRecordTable)

	db.Exec(seeds.UserSeed)
	db.Exec(seeds.TagsSeed)
	db.Exec(seeds.RecordSeed)
}
