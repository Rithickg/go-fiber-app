package routes

import (
	"my-fiber-app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	userRoutes := app.Group("/users")
	userRoutes.Post("/", controllers.CreateUser)

	userRoutes.Get("/", controllers.GetAllUsers)
	userRoutes.Get("/:id", controllers.GetUserByID)
	userRoutes.Put("/:id", controllers.UpdateUser)
	userRoutes.Delete("/:id", controllers.DeleteUser)
}
