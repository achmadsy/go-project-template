package db

import (
	"os"

	"github.com/achmadsy/go-project/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DB = global db var
var DB *gorm.DB

//InitDB mysql
func InitDB() {
	user := os.Getenv("dbuser")
	pass := os.Getenv("dbpass")
	dsn := user + ":" + pass + "@tcp(db:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	defer database.DB()

	if err != nil {
		panic("Failed to connect to database!")
	}
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Address{})

	userDummy := models.User{}
	userDummy.ID = 1

	database.Create(&userDummy)

	DB = database

}
