package room

import (
	"room_app_back/domain/model"
	"room_app_back/domain/repository/room"
)

type IRoomUsecase interface {
	Rooms() (*model.Rooms, error)
	RoomById(id int) (*model.Room, error)
}

type RoomUsecase struct {
	rr room.RoomRepository
}

func NewRoomUsecase(rr room.RoomRepository) *RoomUsecase {
	return &RoomUsecase{rr}
}

func (ru *RoomUsecase) Rooms() (*model.Rooms, error) {
	return ru.rr.FindAll()
}

func (ru *RoomUsecase) RoomById(id int) (*model.Room, error) {
	return ru.rr.FindById(id)
}
