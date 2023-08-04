package database

import "room_app_back/domain/model"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) FindAll() (users model.Users, err error) {
	if err = repo.Find(&users); err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindById(id int) (user model.User, err error) {
	if err = repo.First(&user, id); err != nil {
		return
	}
	return
}

func (repo *UserRepository) Store(u model.User) (user model.User, err error) {
	if err = repo.Create(&u); err != nil {
		return
	}
	user = u
	return
}

func (repo *UserRepository) Update(u model.User) (user model.User, err error) {
	if err = repo.Save(&u); err != nil {
		return
	}
	user = u
	return
}

func (repo *UserRepository) DeleteById(user model.User) (err error) {
	if err = repo.Delete(&user); err != nil {
		return
	}
	return
}
