package model

import "time"

type Users []User

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	LastName  string    `json:"last_name" gorm:"not null"`
	FirstName string    `json:"first_name" gorm:"not null"`
	Age       uint      `json:"age"`
	Rule      string    `json:"rule" gorm:"not null"`
	IdNumber  string    `json:"id_number" gorm:"unique"` // 学籍番号または教員番号
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
