package usecase

import "github.com/just-umyt/blUg/internal/models"

type UserRepo interface {
	Register(user *models.User) error
	Login(user *models.User) error
	Update(id uint, newUser *models.User) error
	Delete(id uint) error
}

type UserUseCaseImpl struct {
	URepo UserRepo
}

func NewUserUsecase(ur UserRepo) *UserUseCaseImpl {
	return &UserUseCaseImpl{URepo: ur}
}

func (ur *UserUseCaseImpl) RegisterUser(user *models.User) error {
	if err := ur.URepo.Register(user); err != nil {
		return err
	}

	return nil
}

func (ur *UserUseCaseImpl) LoginUser(user *models.User) error {
	if err := ur.URepo.Login(user); err != nil {
		return err
	}

	return nil
}

func (ur *UserUseCaseImpl) UpdateUser(id uint, user *models.User) error {
	return ur.URepo.Update(id, user)
}

func (ur *UserUseCaseImpl) DeleteUser(id uint) error {
	return ur.URepo.Delete(id)
}
