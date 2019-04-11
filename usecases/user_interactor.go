package usecases

import "modul1/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	identifier, err := interactor.UserRepository.Store(u)
	if err != nil {
		return
	}
	user, err = interactor.UserRepository.FindById(identifier)
	return
}

func (interactor *UserInteractor) Users() (user domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}

func (interactor *UserInteractor) Login(email string, password string) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByEmail(email, password)
	return
}
