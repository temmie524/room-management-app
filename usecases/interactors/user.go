package interactors

import (
	"room_app_back/domain/model"
	"room_app_back/usecases/port"
)

type UserInteractor struct {
	ur port.UserRepository
}

func NewUserInteractor(userRepository port.UserRepository) *UserInteractor {
	return &UserInteractor{
		ur: userRepository,
	}
}

func (ui *UserInteractor) Add(u model.User) (*model.User, error) {
	return ui.ur.Store(u)
}

func (ui *UserInteractor) Update(u model.User) (*model.User, error) {
	return ui.ur.Update(u)
}

func (ui *UserInteractor) DeleteById(u model.User) error {
	return ui.ur.DeleteById(u)
}

func (ui *UserInteractor) Users() (*model.Users, error) {
	return ui.ur.FindAll()
}

func (ui *UserInteractor) UserById(id int) (*model.User, error) {
	return ui.ur.FindById(id)
}
