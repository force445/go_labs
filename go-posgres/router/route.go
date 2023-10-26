package router

import (
	"go-posgres/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/users", handler.GetAllUser)
	v1.Get("/user/:id", handler.GetUser)
	v1.Post("/user", handler.PostUser)
	v1.Post("/location", handler.PostLocation)
	v1.Get("/locations", handler.GetAllLocation)
	v1.Get("/location/:id", handler.GetLocation)
	v1.Post("/station", handler.PostStation)
	v1.Get("/stations", handler.GetAllStation)
	v1.Get("/station/:id", handler.GetStation)
	v1.Post("/drone", handler.PostDrone)
	v1.Get("/drones", handler.GetAllDrone)
	v1.Get("/drone/:id", handler.GetDrone)
}
