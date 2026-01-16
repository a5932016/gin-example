package databases

import (
	"os"

	"github.com/a5932016/gin_example/databases/postgresql"
	"github.com/a5932016/gin_example/databases/sqlites"
	"github.com/a5932016/gin_example/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	switch os.Getenv("database") {
	case "postgres":
		DB = postgresql.Init()
	default:
		DB = sqlites.Init()
	}

	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		println(err.Error())
	}

	err = seed(DB)
	if err != nil {
		println(err.Error())
	}
}
