package router

import (
	"go-posgres/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	//---------------------------Member--------------------------------
	// v1.Get("/users", handler.GetAllUser)
	// v1.Get("/user/:id", handler.GetUser)
	v1.Post("/user", handler.PostMember)

}
