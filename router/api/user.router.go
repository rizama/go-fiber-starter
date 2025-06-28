package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizama/go-fiber-starter/app/controller"
	"github.com/rizama/go-fiber-starter/app/middleware"
	"github.com/rizama/go-fiber-starter/app/repository"
	"github.com/rizama/go-fiber-starter/app/service"
)

type UserRouter struct {
	app            *fiber.App
	userController controller.UserController
}

func NewUserRouter(app *fiber.App) *UserRouter {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	return &UserRouter{
		app:            app,
		userController: userController,
	}
}

func (r *UserRouter) Setup(api fiber.Router) {
	users := api.Group("/users")
	users.Get("", r.userController.GetUsers)
	users.Get("/:id", r.userController.GetUserByID)
	users.Post("", middleware.DBTransactionHandler(), r.userController.CreateUser)
	users.Put("/:id", middleware.DBTransactionHandler(), r.userController.UpdateUser)
	users.Delete("/:id", middleware.DBTransactionHandler(), r.userController.DeleteUser)
}
