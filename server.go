package main

import (
	"github.com/aminkhn/golang-rest-api/database"
	"github.com/aminkhn/golang-rest-api/routes"

	"github.com/gofiber/fiber/v2"
)

// Handler
func setupRoutes(app *fiber.App) {
	// Welcome EndPoint
	app.Get("/api", welcome)
	// User EndPoint
	app.Group("/api/v1/users")
	app.Post("/api/v1/users", routes.CreateUser)
	app.Get("/api/v1/users", routes.GetUsers)
	app.Get("/api/v1/user/:id")
	app.Put("/api/v1/user/:id")
	app.Delete("/api/v1/user/:id")
}
func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":8000")
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello welcome Here!")
}
