package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/blUg/internal/models"
	"github.com/just-umyt/blUg/internal/token"
)

type UserUsecase interface {
	RegisterUser(user *models.User) error
	LoginUser(user *models.User) error
	UpdateUser(id uint, user *models.User) error
	DeleteUser(id uint) error
}

type UserHandler struct {
	Usercase UserUsecase
}

func NewUserHandler(uu UserUsecase) *UserHandler {
	return &UserHandler{Usercase: uu}
}

func (h *UserHandler) RegisterUserHandler(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		log.Fatal(err)
	}

	// hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	// if err != nil {
	// 	return c.Status(http.StatusBadRequest).SendString("Failed to hash password")
	// }
	if err := h.Usercase.RegisterUser(&user); err != nil {
		log.Fatal("Failed to creat:", err)
		return err
	}

	token, err := token.CreateUserToken(user.ID)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "user_id",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: true,
		// Secure:   true,
	})

	return c.JSON(user)
}

func (h *UserHandler) LoginUserHandler(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		log.Fatal(err)
	}

	if err := h.Usercase.LoginUser(&user); err != nil {
		return err
	}

	token, err := token.CreateUserToken(user.ID)
	if err != nil {
		return err
	}
	c.Cookie(&fiber.Cookie{
		Name:     "user_id",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: true,
		// Secure:   true,
	})


	return c.JSON(user)
}

func (h *UserHandler) UpdateUserHandler(c *fiber.Ctx) error {

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		log.Fatal(err)
	}

	tokenString := c.Cookies("user_id")

	userId, err := token.ParseUserToken(tokenString)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := h.Usercase.UpdateUser(userId, &user); err != nil {
		fmt.Println(err)
		return err
	}

	return c.SendStatus(fiber.StatusOK)

}

func (h *UserHandler) DeleteUserHandler(c *fiber.Ctx) error {

	tokenString := c.Cookies("user_id")

	userId, err := token.ParseUserToken(tokenString)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := h.Usercase.DeleteUser(userId); err != nil {
		log.Fatal("Failed to delete:", err)
	}

	return c.SendStatus(200)
}
