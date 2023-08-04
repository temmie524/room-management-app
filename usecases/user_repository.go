package usecases

import "room_app_back/domain/model"

type UserRepository interface {
	Store(model.User) (model.User, error)
	Update(model.User) (model.User, error)
	DeleteById(model.User) error
	FindAll() (model.Users, error)
	FindById(id int) (model.User, error)
}
