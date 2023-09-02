package router

import (
	"github.com/aminkhn/mysql-rest-api/handlers"
	"github.com/aminkhn/mysql-rest-api/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// API Health Checker
	api := app.Group("/api")
	api.Get("/healthchecker", handlers.Healthchecker)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)
	auth.Post("/logout", handlers.Logout)

	// User
	user := api.Group("/user")
	// Protection
	user.Use(middlewares.Protected())
	user.Use(middlewares.IsBlackListed)
	// User CRUD
	user.Get("/", handlers.GetAllUser)
	user.Get("/:id", handlers.GetUser)
	user.Post("/", handlers.CreateUser)
	user.Put("/:id", handlers.UpdateUserPut)
	user.Patch("/:id", handlers.UpdateUserPatch)
	user.Delete("/:id", handlers.DeleteUser)

	// Product
	//product := api.Group("product")
	//product.Use(middlewares.Protected())
	//product.Use(middlewares.IsBlackListed)
	// Product CRUD
	//product.Get("/", handlers.GetAllProducts)
	//product.Get("/:id", handlers.GetProduct)
	//product.Post("/", handlers.CreateProduct)
	//product.Put("/:id", handlers.UpdateProductPut)
	//product.Patch("/:id", handlers.UpdateProductPatch)
	//product.Delete("/:id", handlers.DeleteProduct)

}
