package user

import (
	"room_app_back/domain/model"
	"room_app_back/domain/repository/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) FindAll() (*model.Users, error) {
	users := model.Users{}

	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil

}

func (repo *UserRepository) FindById(id int) (*model.User, error) {
	user := model.User{}

	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) Store(u *model.User) (*model.User, error) {
	if err := repo.db.Create(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (repo *UserRepository) Update(u *model.User) (*model.User, error) {
	if err := repo.db.Save(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (repo *UserRepository) DeleteById(id int) error {
	user := model.User{}
	if err := repo.db.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := model.User{}
	if err := repo.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil

}
