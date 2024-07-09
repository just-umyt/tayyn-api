package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/blUg/internal/models"
	"github.com/just-umyt/blUg/internal/token"
)

type BlugUsecase interface {
	GetBlug() *[]models.Blug
	GetBlugById(id int) (*models.Blug, error)
	CreateBlug(blug *models.Blug) error
	UpdateBlug(id int, newBlug *models.Blug) error
	DeleteBlug(id int) error
	LikeBlug(blugId int, userId uint) error
}

type BlugHandler struct {
	BUsecase BlugUsecase
}

func NewBlugHandler(bu BlugUsecase) *BlugHandler {
	return &BlugHandler{BUsecase: bu}
}

func (bu *BlugHandler) GetBlugsHandler(c *fiber.Ctx) error {
	blugs := bu.BUsecase.GetBlug()
	return c.JSON(blugs)
}

func (bu *BlugHandler) GetBlugByIdHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	blug, err := bu.BUsecase.GetBlugById(id)
	if err != nil {
		return err
	}

	return c.JSON(blug)
}

func (bu *BlugHandler) CreateBlugHandler(c *fiber.Ctx) error {
	var blug *models.Blug

	if err := c.BodyParser(&blug); err != nil {
		return err
	}

	tokenString := c.Cookies("user_id")
	userId, err := token.ParseUserToken(tokenString)
	if err != nil {
		return err
	}

	blug.UserId = uint(userId)

	if err := bu.BUsecase.CreateBlug(blug); err != nil {
		return err
	}



	return c.JSON(blug)

}

func (bu *BlugHandler) UpdateBlugHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	var newBlug *models.Blug
	if err := c.BodyParser(&newBlug); err != nil {
		return err
	}

	if err := bu.BUsecase.UpdateBlug(id, newBlug); err != nil {
		return err
	}

	return c.JSON(newBlug)
}

func (bu *BlugHandler) DeleteBlugHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	if err := bu.BUsecase.DeleteBlug(id); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (bu *BlugHandler) LikeBlugHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	tokenString := c.Cookies("user_id")

	userId, err := token.ParseUserToken(tokenString)
	if err != nil {
		return err
	}

	if err := bu.BUsecase.LikeBlug(id, userId); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
