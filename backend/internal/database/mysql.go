package database

import (
	"go-vue-app/backend/pkg/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Admin{})
}

func GetMysqlConnection() (*gorm.DB, error) {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	log.Println("dsn: " + dsn)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
