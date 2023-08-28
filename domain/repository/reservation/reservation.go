package reservation

import model "room_app_back/domain/model"

type ReservationRepository interface {
	Store(*model.Reservation) (*model.Reservation, error)
	Update(*model.Reservation) (*model.Reservation, error)
	DeleteById(id int) error
	FindAll() (*model.Reservations, error)
	FindById(id int) (*model.Reservation, error)
}
