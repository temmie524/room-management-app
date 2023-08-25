package reservation

import (
	"room_app_back/domain/model"
	"room_app_back/domain/repository/reservation"
)

type ReservationUsecase struct {
	rr reservation.ReservationRepository
}

func NewReservationUsecase(rr reservation.ReservationRepository) *ReservationUsecase {
	return &ReservationUsecase{
		rr: rr,
	}
}

func (ru *ReservationUsecase) Add(r model.Reservation) (*model.Reservation, error) {
	return ru.rr.Store(r)
}

func (ru *ReservationUsecase) Update(r model.Reservation) (*model.Reservation, error) {
	return ru.rr.Update(r)
}

func (ru *ReservationUsecase) DeleteById(r model.Reservation) error {
	return ru.rr.DeleteById(r)
}

func (ru *ReservationUsecase) Reservations() (*model.Reservations, error) {
	return ru.rr.FindAll()
}

func (ru *ReservationUsecase) ReservationById(id int) (*model.Reservation, error) {
	return ru.rr.FindById(id)
}
