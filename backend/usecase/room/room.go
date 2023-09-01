package room

import (
	"backend/domain/repository/room"
)

type IRoomUsecase interface {
	Rooms() (*AddOutputs, error)
	RoomById(id int) (*AddOutput, error)
}

type RoomUsecase struct {
	rr room.RoomRepository
}

func NewRoomUsecase(rr room.RoomRepository) IRoomUsecase {
	return &RoomUsecase{rr}
}

func (ru *RoomUsecase) Rooms() (*AddOutputs, error) {
	rooms, err := ru.rr.FindAll()
	if err != nil {
		return nil, err
	}
	var outputs AddOutputs
	for _, _r := range *rooms {
		r := _r
		outputs = append(outputs, AddOutput{&r})
	}
	return &outputs, nil

}

func (ru *RoomUsecase) RoomById(id int) (*AddOutput, error) {
	room, err := ru.rr.FindById(id)
	if err != nil {
		return nil, err
	}
	return &AddOutput{room}, nil
}
