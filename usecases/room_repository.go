package usecases

import "room_app_back/domain/model"

type RoomRepository interface {
	FindAll() (model.Rooms, error)
	FindById(id int) (model.Room, error)
}
