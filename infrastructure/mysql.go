package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() *SqlHandler {

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

	SqlHandler := new(SqlHandler)
	SqlHandler.db = db
	return SqlHandler

}

func (handler *SqlHandler) Find(obj interface{}, value ...interface{}) error {
	if err := handler.db.Find(obj).Error; err != nil {
		return err
	}
	return nil
}

func (handler *SqlHandler) First(obj interface{}, where ...interface{}) error {
	if err := handler.db.First(obj, where...).Error; err != nil {
		return err
	}
	return nil
}

// Reservation専用のFind。UserとRoomをPreloadする
func (handler *SqlHandler) FindReservation(obj interface{}, value ...interface{}) error {
	if err := handler.db.Preload("User").Preload("Room").Find(obj).Error; err != nil {
		return err
	}
	return nil
}

func (handler *SqlHandler) Create(obj interface{}) error {
	if err := handler.db.Create(obj).Error; err != nil {
		return err
	}
	return nil
}

func (handler *SqlHandler) Save(obj interface{}) error {
	if err := handler.db.Save(obj).Error; err != nil {
		return err
	}
	return nil
}

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

func (handler *SqlHandler) Raw(sql string, values ...interface{}) error {
	if err := handler.db.Raw(sql, values...).Error; err != nil {
		return err
	}
	return nil
}

/*
func (handler *SqlHandler) Where(sql string, value ...interface{}) error {
	if err := handler.db.Where(sql, value...).Error; err != nil {
		return err
	}
	return nil
}
*/
