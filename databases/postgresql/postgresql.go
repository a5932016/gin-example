package postgresql

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	var err error

	DB, err := gorm.Open(postgres.Open(os.Getenv("portgresStr")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if DB.Error != nil {
		panic(err)
	}

	return DB
}
