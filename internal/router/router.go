package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/blUg/internal/handlers"
	"github.com/just-umyt/blUg/internal/middleware"
	"github.com/just-umyt/blUg/internal/mysqlapi"
	"github.com/just-umyt/blUg/internal/repository"
	"github.com/just-umyt/blUg/internal/usecase"
)

func NewApp() *fiber.App {
	app := fiber.New()

	db := mysqlapi.InitDB()

	userRepo := repository.NewUserRepo(db)
	userUseCases := usecase.NewUserUsecase(userRepo)
	userHandler := handlers.NewUserHandler(userUseCases)

	blugRepo := repository.NewBlugRepo(db)
	blugUseCase := usecase.NewBlugUsecase(blugRepo)
	blugHandler := handlers.NewBlugHandler(blugUseCase)

	middlewareRepo := middleware.NewMiddlewareRepo(db)
	middlewareUsecase := middleware.NewMiddlewareUsecase(middlewareRepo)
	middleareHandler := middleware.NewMiddlewareHandler(middlewareUsecase)

	//users Handler
	app.Post("/register", userHandler.RegisterUserHandler)
	app.Post("/login", userHandler.LoginUserHandler)
	app.Put("/profile", userHandler.UpdateUserHandler)
	app.Delete("/profile", userHandler.DeleteUserHandler)

	//blugs Handler
	app.Get("/blugs", blugHandler.GetBlugsHandler)
	app.Get("/blugs/:id", blugHandler.GetBlugByIdHandler)

	//middleware handler
	midlWare := app.Group("/blugs", middleareHandler.GetUserByIdHandler)

	midlWare.Post("/create", blugHandler.CreateBlugHandler)
	midlWare.Put("/:id", blugHandler.UpdateBlugHandler)
	midlWare.Delete("/:id", blugHandler.DeleteBlugHandler)
	midlWare.Post("/like/:id", blugHandler.LikeBlugHandler)

	return app
}
