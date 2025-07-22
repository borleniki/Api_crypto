package database

import (
	"fmt"
	"kafka-redis/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMySQL() {
	dsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to MySQL!")
	}
	db.AutoMigrate(&models.User{})
	DB = db
	fmt.Println("MySQL connected.")
}
