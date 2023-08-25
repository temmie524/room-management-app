package room

import (
	"room_app_back/domain/model"
	"room_app_back/domain/repository/room"
)

type RoomUsecase struct {
	rr room.RoomRepository
}

func NewRoomUsecase(rr room.RoomRepository) *RoomUsecase {
	return &RoomUsecase{
		rr: rr,
	}
}

func (ru *RoomUsecase) Rooms() (*model.Rooms, error) {
	return ru.rr.FindAll()
}

func (ru *RoomUsecase) RoomById(id int) (*model.Room, error) {
	return ru.rr.FindById(id)
}
