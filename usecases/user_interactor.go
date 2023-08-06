package usecases

import "room_app_back/domain/model"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u model.User) (user model.User, err error) {
	user, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Update(u model.User) (user model.User, err error) {
	user, err = interactor.UserRepository.Update(u)
	return
}

func (interactor *UserInteractor) DeleteById(user model.User) (err error) {
	err = interactor.UserRepository.DeleteById(user)
	return
}

func (interacter *UserInteractor) Users() (users model.Users, err error) {
	users, err = interacter.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (user model.User, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}
