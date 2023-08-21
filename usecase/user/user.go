package user

import (
	"room_app_back/domain/model"
	"room_app_back/infrastructure/repository/user"
)

type UserUsecase struct {
	ur user.UserRepository
}

func NewUserUsecase(userRepository user.UserRepository) *UserUsecase {
	return &UserUsecase{
		ur: userRepository,
	}
}

func (ui *UserUsecase) Add(u model.User) (*model.User, error) {
	return ui.ur.Store(u)
}

func (ui *UserUsecase) Update(u model.User) (*model.User, error) {
	return ui.ur.Update(u)
}

func (ui *UserUsecase) DeleteById(u model.User) error {
	return ui.ur.DeleteById(u)
}

func (ui *UserUsecase) Users() (*model.Users, error) {
	return ui.ur.FindAll()
}

func (ui *UserUsecase) UserById(id int) (*model.User, error) {
	return ui.ur.FindById(id)
}
