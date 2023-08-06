package database

import "room_app_back/domain/model"

type ReservationRepository struct {
	SqlHandler
}

func (repo *ReservationRepository) FindAll() (reservations model.Reservations, err error) {
	if err = repo.FindReservation(&reservations); err != nil {
		return
	}
	return
}

func (repo *ReservationRepository) FindById(id int) (r model.Reservation, err error) {
	if err = repo.First(&r, id); err != nil {
		return
	}
	return
}

func (repo *ReservationRepository) Store(r model.Reservation) (reservation model.Reservation, err error) {
	if err = repo.Create(&r); err != nil {
		return
	}
	reservation = r
	return
}

func (repo *ReservationRepository) Update(r model.Reservation) (reservation model.Reservation, err error) {
	if err = repo.Save(&r); err != nil {
		return
	}
	reservation = r
	return
}

func (repo *ReservationRepository) DeleteById(r model.Reservation) (err error) {
	if err = repo.Delete(&r); err != nil {
		return
	}
	return
}
