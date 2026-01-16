package databases

import (
	"github.com/a5932016/gin_example/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func seed(DB *gorm.DB) error {
	DB.Transaction(func(tx *gorm.DB) error {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

		user := model.User{
			Name:     "test",
			Email:    "test@example.com",
			Password: string(hashPassword),
		}

		if err := tx.First(&model.User{}).Error; err != nil {
			err = tx.Create(&user).Error
			if err != nil {
				return err
			}
		}

		return nil
	})

	return nil
}
