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
	v1.Get("/members", handler.GetAllMember)
	v1.Put("/member/:id", handler.UpdateMember)
	v1.Delete("/member/:id", handler.DeleteMember)

	//---------------------------Creditcards--------------------------------
	v1.Post("/creditcard", handler.PostCreditCard)
	v1.Get("/creditcard/:id", handler.GetCreditCard)
	v1.Put("/creditcard/:id", handler.UpdateCreditCard)
	v1.Delete("/creditcard/:id", handler.DeleteCreditCard)

	//---------------------------Order--------------------------------
	v1.Post("/order", handler.PostOrder)
	v1.Get("/order/:id", handler.GetOrder)
	v1.Get("/orders", handler.GetAllOrder)
	v1.Put("/order/:id", handler.UpdateOrder)
	v1.Delete("/order/:id", handler.DeleteOrder)
}
