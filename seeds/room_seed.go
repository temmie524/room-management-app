package main

import (
	"fmt"
	"log"
	"os"
	"room_app_back/domain/model"
	"time"

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
	if err := RoomSeeds(db); err != nil {
		fmt.Println(err)
	}
	defer fmt.Println("Successfully migrated seeds!!")
	defer sqlDB.Close()

	if err != nil {
		log.Fatalln(err)
	}
}

func RoomSeeds(db *gorm.DB) error {
	room1 := model.Room{
		ID:        1,
		RoomNum:   414,
		Building:  "A",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&room1).Error; err != nil {
		fmt.Println(err)
	}

	room2 := model.Room{
		ID:        2,
		RoomNum:   301,
		Building:  "M",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&room2).Error; err != nil {
		fmt.Println(err)
	}

	room3 := model.Room{
		ID:        3,
		RoomNum:   201,
		Building:  "M",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&room3).Error; err != nil {
		fmt.Println(err)
	}

	return nil
}
