package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if os.Getenv("GO_ENV") == "production" {
		dbUser = os.Getenv("AWS_DB_USER")
		dbPass = os.Getenv("AWS_DB_PASS")
		dbHost = os.Getenv("AWS_DB_HOST")
		dbPort = os.Getenv("AWS_DB_PORT")
		dbName = os.Getenv("AWS_DB_NAME")

	}

	if os.Getenv("GO_ENV") == "docker" {
		dbHost = os.Getenv("DOCKER_DB_HOST")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	return db

}
