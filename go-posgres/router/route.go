package router

import (
	"go-posgres/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	//---------------------------Member--------------------------------
	v1.Post("/member", handler.PostMember)
	v1.Get("/member/:id", handler.GetMember)
	v1.Get("/member", handler.GetAllMember)
	v1.Put("/member/:id", handler.UpdateMember)
	v1.Delete("/member/:id", handler.DeleteMember)

}
