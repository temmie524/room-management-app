package usecases

import "room_app_back/domain/model"

type RoomInteractor struct {
	RoomRepository RoomRepository
}

func (interactor *RoomInteractor) Rooms() (rooms model.Rooms, err error) {
	rooms, err = interactor.RoomRepository.FindAll()
	return
}

func (interactor *RoomInteractor) RoomById(id int) (room model.Room, err error) {
	room, err = interactor.RoomRepository.FindById(id)
	return
}
