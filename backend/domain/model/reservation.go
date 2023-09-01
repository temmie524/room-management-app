package model

import "time"

type Reservations []Reservation

type Reservation struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Purpose   string    `json:"purpose"`
	StartTime string    `json:"start_time" gorm:"not null"`
	EndTime   string    `json:"end_time" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	RoomID    uint      `json:"room_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User User
	Room Room
}
