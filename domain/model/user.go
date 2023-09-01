package model

import "time"

type Users []User

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	LastName  string    `json:"last_name" gorm:"not null"`
	FirstName string    `json:"first_name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Age       uint      `json:"age"`
	Role      string    `json:"role"`
	IdNumber  string    `json:"id_number" gorm:"unique"` // 学籍番号または教員番号
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Email string `json:"email" gorm:"unique"`
}
