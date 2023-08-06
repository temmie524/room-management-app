package model

import "time"

type Room struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	RoomNum   uint      `json:"room_num" gorm:"not null"`
	Building  string    `json:"building" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Reservations []Reservation `gorm:"ForeignKey:RoomID"`
}
