package database

import "room_app_back/domain/model"

type RoomRepository struct {
	SqlHandler
}

func (repo *RoomRepository) FindAll() (rooms model.Rooms, err error) {
	if err = repo.Find(&rooms); err != nil {
		return
	}
	return
}

func (repo *RoomRepository) FindById(id int) (room model.Room, err error) {
	if err = repo.First(&room, id); err != nil {
		return
	}
	return
}
