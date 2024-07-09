package repository

import (
	"errors"

	"github.com/just-umyt/blUg/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	Database *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{Database: db}
}

func (repo *UserRepository) Register(user *models.User) error {
	if err := repo.Database.Create(&user).Error; err != nil {
		return nil
	}

	return nil
}

func (repo *UserRepository) Login(user *models.User) error {
	var loginUser models.User
	repo.Database.First(&loginUser, "email = ?", user.Email)
	if loginUser.ID == 0 {
		return errors.New("invalid email")
	}

	// err := bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(user.Password))
	if loginUser.Password != user.Password {
		return errors.New("wrong password")
	}

	*user = loginUser

	return nil

}

func (repo *UserRepository) Update(id uint, newUser *models.User) error {

	var oldUser models.User

	if err := repo.Database.First(&oldUser, id).Error; err != nil {
		return err
	}

	oldUser.Name = newUser.Name
	oldUser.Email = newUser.Email
	oldUser.Nick = newUser.Nick

	repo.Database.Save(&oldUser)

	return nil
}

func (repo *UserRepository) Delete(id uint) error {

	if err := repo.Database.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

// func (repo *UserRepository) GetUser(id float64) error {

// }
