package usecases

import "room_app_back/domain/model"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u model.User) (model.User, error) {
	return interactor.UserRepository.Store(u)
}

func (interactor *UserInteractor) Update(u model.User) (model.User, error) {
	return interactor.UserRepository.Update(u)
}

func (interactor *UserInteractor) DeleteById(u model.User) error {
	return interactor.UserRepository.DeleteById(u)
}

func (interactor *UserInteractor) Users() (model.Users, error) {
	return interactor.UserRepository.FindAll()
}

func (interactor *UserInteractor) UserById(identifier int) (model.User, error) {
	return interactor.UserRepository.FindById(identifier)
}
