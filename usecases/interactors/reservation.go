package interactors

import (
	"room_app_back/domain/model"
	"room_app_back/usecases/port"
)

type ReservationInteractor struct {
	rr port.ReservationRepository
}

func NewReservationInteractor(rr port.ReservationRepository) *ReservationInteractor {
	return &ReservationInteractor{
		rr: rr,
	}
}

func (ri *ReservationInteractor) Add(r model.Reservation) (*model.Reservation, error) {
	return ri.rr.Store(r)
}

func (ri *ReservationInteractor) Update(r model.Reservation) (*model.Reservation, error) {
	return ri.rr.Update(r)
}

func (ri *ReservationInteractor) DeleteById(r model.Reservation) error {
	return ri.rr.DeleteById(r)
}

func (ri *ReservationInteractor) Reservations() (*model.Reservations, error) {
	return ri.rr.FindAll()
}

func (ri *ReservationInteractor) ReservationById(identifier int) (*model.Reservation, error) {
	return ri.rr.FindById(identifier)
}
