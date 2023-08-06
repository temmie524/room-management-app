package usecases

import "room_app_back/domain/model"

type ReservationInteractor struct {
	ReservationRepository ReservationRepository
}

func (interactor *ReservationInteractor) Add(r model.Reservation) (reservation model.Reservation, err error) {
	reservation, err = interactor.ReservationRepository.Store(r)
	return
}

func (interactor *ReservationInteractor) Update(r model.Reservation) (reservation model.Reservation, err error) {
	reservation, err = interactor.ReservationRepository.Update(r)
	return
}

func (interactor *ReservationInteractor) DeleteById(r model.Reservation) (err error) {
	err = interactor.ReservationRepository.DeleteById(r)
	return
}

func (interacter *ReservationInteractor) Reservations() (reservations model.Reservations, err error) {
	reservations, err = interacter.ReservationRepository.FindAll()
	return
}

func (interactor *ReservationInteractor) ReservationById(identifier int) (reservation model.Reservation, err error) {
	reservation, err = interactor.ReservationRepository.FindById(identifier)
	return
}
