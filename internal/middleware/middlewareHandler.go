package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/blUg/internal/models"
)

type MiddlewareUsecase interface {
	GetUserById(userId uint) (*models.User, error)
}

type MiddlewareHandler struct {
	MUsecase MiddlewareUsecase
}

func NewMiddlewareHandler(mu MiddlewareUsecase) *MiddlewareHandler {
	return &MiddlewareHandler{MUsecase: mu}
}

func (mh *MiddlewareHandler) GetUserByIdHandler(c *fiber.Ctx) error {

	tokenString := c.Cookies("user_id")
	if tokenString == "" {
		fmt.Println("user not logged in")
		return c.Redirect("/login")
	}

	fmt.Println("user logged in")
	return c.Next()
}
