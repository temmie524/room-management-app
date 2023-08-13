package interactors

import (
	"room_app_back/domain/model"
	"room_app_back/usecases/port"
)

type RoomInteractor struct {
	rr port.RoomRepository
}

func NewRoomInteractor(rr port.RoomRepository) *RoomInteractor {
	return &RoomInteractor{
		rr: rr,
	}
}

func (i *RoomInteractor) Rooms() (*model.Rooms, error) {
	return i.rr.FindAll()
}

func (i *RoomInteractor) RoomById(id int) (*model.Room, error) {
	return i.rr.FindById(id)
}
