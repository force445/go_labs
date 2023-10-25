package main

import (
	databases "labs/database"
	"labs/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	databases.ConnectMongoDB()

	routes.AuthorRoute(app)

	app.Listen(":6000")

}
