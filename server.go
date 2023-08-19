package main

import (
	"github.com/aminkhn/golang-rest-api/database"
	"github.com/aminkhn/golang-rest-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// establishing connection to database + migrations
	database.MysqlConnectDb()
	// new instance of fiber
	app := fiber.New()
	// setting up URIs routes
	router.SetupRoutes(app)
	// staring webserver
	app.Listen(":8000")
}
