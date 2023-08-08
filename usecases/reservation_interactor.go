package usecases

import "room_app_back/domain/model"

type ReservationInteractor struct {
	ReservationRepository ReservationRepository
}

func (i *ReservationInteractor) Add(r model.Reservation) (model.Reservation, error) {
	return i.ReservationRepository.Store(r)
}

func (i *ReservationInteractor) Update(r model.Reservation) (model.Reservation, error) {
	return i.ReservationRepository.Update(r)
}

func (i *ReservationInteractor) DeleteById(r model.Reservation) error {
	return i.ReservationRepository.DeleteById(r)
}

func (interacter *ReservationInteractor) Reservations() (model.Reservations, error) {
	return interacter.ReservationRepository.FindAll()
}

func (i *ReservationInteractor) ReservationById(identifier int) (model.Reservation, error) {
	return i.ReservationRepository.FindById(identifier)

}
