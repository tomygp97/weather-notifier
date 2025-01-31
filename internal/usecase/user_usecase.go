package usecase

import (
	"errors"

	"github.com/tomygp97/weather-notifier/internal/domain"
)

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{repo}
}

func (u *UserUsecase) RegisterUser(user *domain.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("el nombre y el email son obligatorios")
	}

	if user.OptedOut == nil {
		defaultOptOut := false
		user.OptedOut = &defaultOptOut
	}

	return u.repo.Save(user)
}

func (u *UserUsecase) GetUsers() ([]domain.User, error) {
	return u.repo.FindAll()
}

func (u *UserUsecase) GetSingleUser(id int) (*domain.User, error) {
	return u.repo.FindByID(id)
}

func (u *UserUsecase) UpdateUser(user *domain.User) error {
	return u.repo.Update(user)
}

func (u *UserUsecase) DeleteUser(id int) error {
	return u.repo.Delete(id)
}
