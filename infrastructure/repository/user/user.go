package user

import (
	"room_app_back/domain/model"
)

type UserRepository struct {
	SqlHandler
}

func NewUserRepository(sqlHandler SqlHandler) *UserRepository {
	return &UserRepository{
		SqlHandler: sqlHandler,
	}
}

func (repo *UserRepository) FindAll() (*model.Users, error) {
	var users model.Users

	if err := repo.Find(&users); err != nil {
		return nil, err
	}
	return &users, nil

}

func (repo *UserRepository) FindById(id int) (*model.User, error) {
	var user model.User

	if err := repo.First(&user, id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) Store(u model.User) (*model.User, error) {
	if err := repo.Create(&u); err != nil {
		return nil, err
	}
	var user model.User = u
	return &user, nil
}

func (repo *UserRepository) Update(u model.User) (*model.User, error) {
	if err := repo.Save(&u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (repo *UserRepository) DeleteById(user model.User) error {
	if err := repo.Delete(&user); err != nil {
		return err
	}
	return nil
}

// 引数については改修するかも
func (repo *UserRepository) FindByEmail(u *model.User, email string) error {
	if err := repo.First(u, "email=?", email); err != nil {
		return err
	}
	return nil

}
