package database

import (
	"backend-assignment/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	database, err := gorm.Open(
		sqlite.Open("test.db"),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)

	}
	DB = database

}
func CreateUserTable() {
	DB.AutoMigrate(&models.User{})
}
