package router

import (
	"github.com/aminkhn/golang-rest-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Welcome EndPoint
	app.Get("/api", handlers.Welcome)
	// ..... User EndPoints
	u := app.Group("/api/v1/users")
	u.Post("/", handlers.CreateUser)
	u.Get("/", handlers.GetUsers)
	u.Get("/:id", handlers.GetUser)
	//u.Put("/:id")
	//u.Delete("/:id")
	// ..... Product EndPoints
	/*
		p := app.Group("/api/v1/products")
		p.Post("/", handlers.CreateUser)
		p.Get("/", handlers.GetUsers)
		p.Get("/:id", handlers.GetUser)
		p.Put("/:id")
		p.Delete("/:id")
	*/
	// ..... Order Endpoints
	/*
		o := app.Group("/api/v1/orders")
		o.Post("/", handlers.CreateUser)
		o.Get("/", handlers.GetUsers)
		o.Get("/:id", handlers.GetUser)
		o.Put("/:id")
		o.Delete("/:id")
	*/
}
