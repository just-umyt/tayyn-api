package repository

import (
	"github.com/just-umyt/blUg/internal/models"
	"gorm.io/gorm"
)

type BlugRepository struct {
	Database *gorm.DB
}

func NewBlugRepo(db *gorm.DB) *BlugRepository {
	return &BlugRepository{Database: db}
}

func (repo *BlugRepository) Get() *[]models.Blug {
	var blugs []models.Blug
	repo.Database.Find(&blugs)
	return &blugs
}

func (repo *BlugRepository) GetById(id int) (*models.Blug, error) {

	var blug models.Blug

	if err := repo.Database.First(&blug, id).Error; err != nil {
		return &blug, err
	}

	return &blug, nil
}

func (repo *BlugRepository) Create(blug *models.Blug) error {
	if err := repo.Database.Create(&blug).Error; err != nil {
		return err
	}
	return nil
}

func (repo *BlugRepository) Update(id int, newBlug *models.Blug) error {
	var oldBlug models.Blug

	repo.Database.First(&oldBlug, id)

	oldBlug.Title = newBlug.Title
	oldBlug.Content = newBlug.Content

	if err := repo.Database.Save(&oldBlug).Error; err != nil {
		return err
	}
	return nil
}

func (repo *BlugRepository) Delete(id int) error {

	if err := repo.Database.Delete(&models.Blug{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (repo *BlugRepository) Like(blugId int, userId uint) error {
	var user models.User
	repo.Database.First(&user, userId)

	var blug models.Blug
	repo.Database.First(&blug, blugId)

	var users []models.User
	repo.Database.Model(&blug).Association("Likes").Find(&users)

	for _, u := range users {
		if u.ID == uint(userId) {
			return repo.Database.Model(&blug).Association("Likes").Delete(&user)
		}
	}

	return repo.Database.Model(&blug).Association("Likes").Append(&user)
}
