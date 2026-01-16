package sqlites

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	var err error

	dbPath := "./databases/sqlites/oprueba.db"

	DB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if DB.Error != nil {
		panic(err)
	}

	return DB
}
