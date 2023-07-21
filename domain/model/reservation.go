package model

import "time"

type Reservation struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Purpose   string    `json:"purpose"`
	StartTime string    `json:"start_time" gorm:"not null"`
	EndTime   string    `json:"end_time"`
	UserId    uint      `json:"user_id"`
	RoomId    uint      `json:"room_id" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 以下は外部キー
	User User `gorm:"foreignkey:UserId"`
	Room Room `gorm:"foreignkey:RoomId"`
}
