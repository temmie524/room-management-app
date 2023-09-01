package room

import (
	"backend/domain/model"
	"backend/domain/repository/room"

	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) room.RoomRepository {
	return &RoomRepository{db}
}

func (repo *RoomRepository) FindAll() (*model.Rooms, error) {
	rooms := model.Rooms{}
	if err := repo.db.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return &rooms, nil
}

func (repo *RoomRepository) FindById(id int) (*model.Room, error) {
	room := model.Room{}
	if err := repo.db.First(&room, id).Error; err != nil {
		return nil, err
	}
	return &room, nil
}
