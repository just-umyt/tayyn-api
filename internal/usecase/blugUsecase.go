package usecase

import "github.com/just-umyt/blUg/internal/models"

type BlugRepo interface {
	Get() *[]models.Blug
	GetById(id int) (*models.Blug, error)
	Create(blug *models.Blug) error
	Update(id int, newBlug *models.Blug) error
	Delete(id int) error
	Like(blugId int, userId uint) error
}

type BlugUsecaseImpl struct {
	BRepo BlugRepo
}

func NewBlugUsecase(br BlugRepo) *BlugUsecaseImpl {
	return &BlugUsecaseImpl{BRepo: br}
}

func (bl *BlugUsecaseImpl) GetBlug() *[]models.Blug {
	return bl.BRepo.Get()
}

func (bl *BlugUsecaseImpl) GetBlugById(id int) (*models.Blug, error) {
	return bl.BRepo.GetById(id)
}

func (bl *BlugUsecaseImpl) CreateBlug(blug *models.Blug) error {
	return bl.BRepo.Create(blug)
}

func (bl *BlugUsecaseImpl) UpdateBlug(id int, newBlug *models.Blug) error {
	return bl.BRepo.Update(id, newBlug)
}

func (bl *BlugUsecaseImpl) DeleteBlug(id int) error {
	return bl.BRepo.Delete(id)
}

func (bl *BlugUsecaseImpl) LikeBlug(blugId int, userId uint) error {
	return bl.BRepo.Like(blugId, userId)
}
