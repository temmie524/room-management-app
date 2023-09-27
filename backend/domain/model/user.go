package model

import "time"

type Users []User

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	LastName  string    `json:"last_name"`
	FirstName string    `json:"first_name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Age       uint      `json:"age"`
	Role      string    `json:"role"`
	IdNumber  string    `json:"id_number"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
