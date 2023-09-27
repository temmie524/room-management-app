package user

import "time"

type AddInputs []AddInputs

type AddInput struct {
	ID        uint      `json:"id"`
	LastName  string    `json:"last_name"`
	FirstName string    `json:"first_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       uint      `json:"age"`
	Role      string    `json:"role"`
	IdNumber  string    `json:"id_number"` // 学籍番号または教員番号
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
