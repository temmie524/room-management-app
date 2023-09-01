package main

import (
	"backend/domain/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalln(err)
		}
	}
	//以下の処理は分割するかも
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, _ := db.DB()
	defer fmt.Println("Successfully migrated database!!")
	defer sqlDB.Close()

	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Debug().AutoMigrate(&model.User{}, &model.Room{}, &model.Reservation{}); err != nil {
		fmt.Println("Database migrate failed.")
	}
}
