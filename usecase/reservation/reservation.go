package reservation

import (
	"room_app_back/domain/model"
	"room_app_back/domain/repository/reservation"
)

type IReservationUsecase interface {
	Add(r *model.Reservation) (*model.Reservation, error)
	Update(r *model.Reservation) (*model.Reservation, error)
	DeleteById(id int) error
	Reservations() (*model.Reservations, error)
	ReservationById(id int) (*model.Reservation, error)
}

type ReservationUsecase struct {
	rr reservation.ReservationRepository
}

func NewReservationUsecase(rr reservation.ReservationRepository) *ReservationUsecase {
	return &ReservationUsecase{rr}
}

func (ru *ReservationUsecase) Add(r *model.Reservation) (*model.Reservation, error) {
	return ru.rr.Store(r)
}

func (ru *ReservationUsecase) Update(r *model.Reservation) (*model.Reservation, error) {
	return ru.rr.Update(r)
}

func (ru *ReservationUsecase) DeleteById(id int) error {
	return ru.rr.DeleteById(id)
}

func (ru *ReservationUsecase) Reservations() (*model.Reservations, error) {
	return ru.rr.FindAll()
}

func (ru *ReservationUsecase) ReservationById(id int) (*model.Reservation, error) {
	return ru.rr.FindById(id)
}
