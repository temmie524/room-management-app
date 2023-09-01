package room

import "time"

type Rooms []Room

type Room struct {
	ID        uint
	RoomNum   uint
	Building  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
