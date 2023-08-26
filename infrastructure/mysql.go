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

	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db

}

/*
func (handler *SqlHandler) Delete(obj interface{}, value ...interface{}) error {
	if err := handler.db.Delete(obj).Error; err != nil {
		return err
	}
	return nil
}

func (handler *SqlHandler) Exec(sql string, values ...interface{}) error {
	if err := handler.db.Exec(sql, values...).Error; err != nil {
		return err
	}
	return nil
}

func (handler *SqlHandler) FindReservation(obj interface{}, value ...interface{}) error {
	if err := handler.db.Preload("User").Preload("Room").Find(obj).Error; err != nil {
		return err
	}
	return nil
}
*/
