package reservation

import "room_app_back/domain/model"

type ReservationRepository struct {
	SqlHandler
}

func NewReservationRepository(sqlHandler SqlHandler) *ReservationRepository {
	return &ReservationRepository{
		SqlHandler: sqlHandler,
	}
}

func (repo *ReservationRepository) FindAll() (*model.Reservations, error) {
	rs := model.Reservations{}

	if err := repo.FindReservation(&rs); err != nil {
		return nil, err
	}
	return &rs, nil
}

func (repo *ReservationRepository) FindById(id int) (*model.Reservation, error) {
	r := model.Reservation{}

	if err := repo.First(&r, id); err != nil {
		return nil, err
	}
	return &r, nil
}

func (repo *ReservationRepository) Store(r model.Reservation) (*model.Reservation, error) {
	if err := repo.Create(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (repo *ReservationRepository) Update(r model.Reservation) (*model.Reservation, error) {
	if err := repo.Save(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (repo *ReservationRepository) DeleteById(r model.Reservation) error {
	if err := repo.Delete(&r); err != nil {
		return err
	}
	return nil
}
