package middleware

import "github.com/just-umyt/blUg/internal/models"

type MiddlewareRepo interface {
	GetById(userId uint) (*models.User, error)
}

type MiddlewareUsecaseImpl struct {
	MRepo MiddlewareRepo
}

func NewMiddlewareUsecase(mr MiddlewareRepo) *MiddlewareUsecaseImpl {
	return &MiddlewareUsecaseImpl{MRepo: mr}
}

func (mu *MiddlewareUsecaseImpl) GetUserById(userId uint) (*models.User, error) {
	return mu.MRepo.GetById(userId)
}
