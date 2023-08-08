package usecases

import "room_app_back/domain/model"

type RoomInteractor struct {
	RoomRepository RoomRepository
}

func (i *RoomInteractor) Rooms() (model.Rooms, error) {
	return i.RoomRepository.FindAll()
}

func (i *RoomInteractor) RoomById(id int) (model.Room, error) {
	return i.RoomRepository.FindById(id)
}
