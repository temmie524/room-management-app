package room

import "backend/domain/model"

type RoomRepository interface {
	FindAll() (*model.Rooms, error)
	FindById(id int) (*model.Room, error)
}
