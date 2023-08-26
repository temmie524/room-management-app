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
	var users model.Users

	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil

}

func (repo *UserRepository) FindById(id int) (*model.User, error) {
	var user model.User

	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) Store(u model.User) (*model.User, error) {
	if err := repo.db.Create(&u).Error; err != nil {
		return nil, err
	}
	var user model.User = u
	return &user, nil
}

func (repo *UserRepository) Update(u model.User) (*model.User, error) {
	if err := repo.db.Save(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (repo *UserRepository) DeleteById(user model.User) error {
	if err := repo.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

// 引数については改修するかも
func (repo *UserRepository) FindByEmail(u *model.User, email string) error {
	if err := repo.db.First(u, "email=?").Error; err != nil {
		return err
	}
	return nil

}
