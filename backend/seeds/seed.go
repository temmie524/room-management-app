package main

import (
	"backend/domain/model"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}
	//以下の処理は分割するかも
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, _ := db.DB()
	if err := roomSeeds(db); err != nil {
		fmt.Println(err)
	}
	if err := userSeeds(db); err != nil {
		fmt.Println(err)
	}
	defer fmt.Println("Successfully migrated seeds!!")
	defer sqlDB.Close()

	if err != nil {
		log.Fatalln(err)
	}
}

func roomSeeds(db *gorm.DB) error {
	rooms := []model.Room{
		{
			ID:        1,
			RoomNum:   414,
			Building:  "A",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			RoomNum:   301,
			Building:  "M",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        3,
			RoomNum:   201,
			Building:  "M",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	for _, _room := range rooms {
		room := _room
		if err := db.Create(&room).Error; err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

func userSeeds(db *gorm.DB) error {
	// Insert sample user data
	users := []model.User{
		{
			LastName:  "Doe",
			FirstName: "John",
			Email:     "john@example.com",
			Password:  "hashed_password",
			Age:       28,
			Role:      "user",
			IdNumber:  "123456789",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			LastName:  "Smith",
			FirstName: "Jane",
			Email:     "jane@example.com",
			Password:  "hashed_password",
			Age:       24,
			Role:      "user",
			IdNumber:  "987654321",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, _user := range users {
		user := _user
		if err := db.Create(&user).Error; err != nil {
			fmt.Println("Error creating user:", err)
		}
	}

	return nil
}
