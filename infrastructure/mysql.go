package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type sqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() *sqlHandler {

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

	sqlHandler := new(sqlHandler)
	sqlHandler.db = db
	return sqlHandler

}

/*
func (handler *sqlHandler) Find(obj interface{}, value ...interface{}) error {
	if err := handler.db.Find(obj).Error; err != nil {
		return err
	}
	return nil
}
*/

func (handler *sqlHandler) First(obj interface{}, where ...interface{}) error {
	if err := handler.db.First(obj, where...).Error; err != nil {
		return err
	}
	return nil
}

// Reservation専用のFind。UserとRoomをPreloadする
func (handler *sqlHandler) Find(obj interface{}, value ...interface{}) error {
	if err := handler.db.Preload("User").Preload("Room").Find(obj).Error; err != nil {
		return err
	}
	return nil
}

func (handler *sqlHandler) Create(obj interface{}) error {
	if err := handler.db.Create(obj).Error; err != nil {
		return err
	}
	return nil
}

func (handler *sqlHandler) Save(obj interface{}) error {
	if err := handler.db.Save(obj).Error; err != nil {
		return err
	}
	return nil
}

func (handler *sqlHandler) Delete(obj interface{}, value ...interface{}) error {
	if err := handler.db.Delete(obj).Error; err != nil {
		return err
	}
	return nil
}

func (handler *sqlHandler) Exec(sql string, values ...interface{}) error {
	if err := handler.db.Exec(sql, values...).Error; err != nil {
		return err
	}
	return nil
}

func (handler *sqlHandler) Raw(sql string, values ...interface{}) error {
	if err := handler.db.Raw(sql, values...).Error; err != nil {
		return err
	}
	return nil
}
