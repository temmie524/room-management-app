package room

import "room_app_back/domain/model"

type RoomRepository struct {
	SqlHandler
}

func NewRoomRepository(sqlHandler SqlHandler) *RoomRepository {
	return &RoomRepository{
		SqlHandler: sqlHandler,
	}
}

func (repo *RoomRepository) FindAll() (*model.Rooms, error) {
	rooms := model.Rooms{}
	if err := repo.Find(&rooms); err != nil {
		return nil, err
	}
	return &rooms, nil
}

func (repo *RoomRepository) FindById(id int) (*model.Room, error) {
	room := model.Room{}
	if err := repo.First(&room, id); err != nil {
		return nil, err
	}
	return &room, nil
}
