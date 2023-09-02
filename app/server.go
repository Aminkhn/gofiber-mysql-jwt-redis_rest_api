package app

import (
	"log"

	"github.com/aminkhn/mysql-rest-api/config"
	"github.com/aminkhn/mysql-rest-api/database"
	"github.com/aminkhn/mysql-rest-api/router"
	"github.com/gofiber/fiber/v2"
)

func Server() {
	// loading Env variables
	loadConfig, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("can not load Envirnment variables", err)
	}
	// establishing connection to database + migrations
	database.MysqlConnectDb(&loadConfig)
	// connecting to Redis
	database.RedisConnectDb(&loadConfig)
	// new instance of fiber
	app := fiber.New()
	// setting up URIs routes
	router.SetupRoutes(app)
	// staring webserver
	log.Fatal(app.Listen(":" + loadConfig.SERVER_Port))
}
