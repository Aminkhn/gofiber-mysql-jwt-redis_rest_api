package router

import (
	"github.com/aminkhn/golang-rest-api/handlers"
	"github.com/aminkhn/golang-rest-api/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handlers.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)

	// User
	user := api.Group("/user")
	user.Get("/:id", handlers.GetUser)
	user.Post("/", handlers.CreateUser)
	user.Patch("/:id", middlewares.Protected(), handlers.UpdateUser)
	user.Delete("/:id", middlewares.Protected(), handlers.DeleteUser)

	// Product
	product := api.Group("/product")
	product.Get("/", handlers.GetAllProducts)
	product.Get("/:id", handlers.GetProduct)
	product.Post("/", middlewares.Protected(), handlers.CreateProduct)
	product.Delete("/:id", middlewares.Protected(), handlers.DeleteProduct)
	// Order
	/*
		order := api.Group("/order")
		order.Get("/", handlers.GetAllOrders)
		order.Get("/:id", handlers.GetOrder)
		order.Post("/", middlewares.Protected(), handlers.CreateOrder)
		order.Delete("/:id", middlewares.Protected(), handlers.DeleteOrder)
	*/
}
