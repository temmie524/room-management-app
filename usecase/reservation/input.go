package reservation

import (
	"backend/domain/model"
	"time"
)

type AddInputs []AddInput

type AddInput struct {
	ID        uint
	Purpose   string
	StartTime string
	EndTime   string
	UserID    uint
	RoomID    uint
	CreatedAt time.Time
	UpdatedAt time.Time

	User model.User
	Room model.Room
}
