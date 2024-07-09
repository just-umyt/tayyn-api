package middleware

import (
	"errors"

	"github.com/just-umyt/blUg/internal/models"
	"gorm.io/gorm"
)

type MiddlewareRepoImpl struct {
	Database *gorm.DB
}

func NewMiddlewareRepo(db *gorm.DB) *MiddlewareRepoImpl {
	return &MiddlewareRepoImpl{Database: db}
}

func (mr *MiddlewareRepoImpl) GetById(userId uint) (*models.User, error) {
	var user *models.User
	mr.Database.First(&user, userId)
	if user.ID == 0 {
		return &models.User{}, errors.New("user not logged in")
	}

	return user, nil
}
