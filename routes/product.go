package routes

import (
	"my-fiber-app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(app *fiber.App) {
	productRoutes := app.Group("/products")

	productRoutes.Get("/", controllers.GetAllProducts)
	productRoutes.Get("/:id", controllers.GetProductByID)
	productRoutes.Post("/", controllers.CreateProduct)
	productRoutes.Put("/:id", controllers.UpdateProduct)
	productRoutes.Delete("/:id", controllers.DeleteProduct)
}
