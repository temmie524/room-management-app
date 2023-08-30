package reservation

import (
	"backend/domain/model"
	"backend/domain/repository/reservation"

	"gorm.io/gorm"
)

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) reservation.ReservationRepository {
	return &ReservationRepository{db}
}

func (repo *ReservationRepository) FindAll() (*model.Reservations, error) {
	rs := model.Reservations{}

	if err := repo.db.Find(&rs).Error; err != nil {
		return nil, err
	}
	return &rs, nil
}

func (repo *ReservationRepository) FindById(id int) (*model.Reservation, error) {
	r := model.Reservation{}

	if err := repo.db.First(&r, id).Error; err != nil {
		return nil, err
	}
	return &r, nil
}

func (repo *ReservationRepository) Store(r *model.Reservation) (*model.Reservation, error) {
	if err := repo.db.Create(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func (repo *ReservationRepository) Update(r *model.Reservation) (*model.Reservation, error) {
	if err := repo.db.Save(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func (repo *ReservationRepository) DeleteById(id int) error {
	r := model.Reservation{}
	if err := repo.db.Delete(&r).Error; err != nil {
		return err
	}
	return nil
}
