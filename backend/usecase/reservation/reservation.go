package reservation

import (
	"backend/domain/model"
	"backend/domain/repository/reservation"
)

type IReservationUsecase interface {
	Add(input *AddInput) (*AddOutput, error)
	Update(input *AddInput) (*AddOutput, error)
	DeleteById(id int) error
	Reservations() (*AddOutputs, error)
	ReservationById(id int) (*AddOutput, error)
}

type ReservationUsecase struct {
	rr reservation.ReservationRepository
}

func NewReservationUsecase(rr reservation.ReservationRepository) IReservationUsecase {
	return &ReservationUsecase{rr}
}

func (ru *ReservationUsecase) Add(input *AddInput) (*AddOutput, error) {
	output, err := ru.rr.Store(&model.Reservation{
		ID:        input.ID,
		Purpose:   input.Purpose,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
		UserID:    input.UserID,
		RoomID:    input.RoomID,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		User:      input.User,
		Room:      input.Room,
	})
	if err != nil {
		return nil, err
	}
	return &AddOutput{output}, nil
}

func (ru *ReservationUsecase) Update(input *AddInput) (*AddOutput, error) {
	output, err := ru.rr.Update(&model.Reservation{
		ID:        input.ID,
		Purpose:   input.Purpose,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
		UserID:    input.UserID,
		RoomID:    input.RoomID,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		User:      input.User,
		Room:      input.Room,
	})
	if err != nil {
		return nil, err
	}
	return &AddOutput{output}, nil
}

func (ru *ReservationUsecase) DeleteById(id int) error {
	return ru.rr.DeleteById(id)
}

func (ru *ReservationUsecase) Reservations() (*AddOutputs, error) {
	users, err := ru.rr.FindAll()
	if err != nil {
		return nil, err
	}
	var outputs AddOutputs
	for _, _u := range *users {
		u := _u
		outputs = append(outputs, AddOutput{&u})
	}
	return &outputs, nil
}

func (ru *ReservationUsecase) ReservationById(id int) (*AddOutput, error) {
	output, err := ru.rr.FindById(id)
	if err != nil {
		return nil, err
	}
	return &AddOutput{output}, nil
}
