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

	//---------------------------Book--------------------------------
	v1.Post("/book", handler.PostBook)
	v1.Get("/book/:id", handler.GetBook)
	v1.Get("/books", handler.GetAllBook)
	v1.Put("/book/:id", handler.UpdateBook)
	v1.Delete("/book/:id", handler.DeleteBook)

	//---------------------------Author--------------------------------
	v1.Post("/author", handler.PostAuthor)
	v1.Get("/author/:id", handler.GetAuthor)
	v1.Put("/author/:id", handler.UpdateAuthor)
	v1.Delete("/author/:id", handler.DeleteAuthor)

	//---------------------------Author--------------------------------
	v1.Post("/publisher", handler.PostPublisher)
	v1.Get("/publisher/:id", handler.GetPublisher)
	v1.Put("/publisher/:id", handler.UpdatePublisher)
	v1.Delete("/publisher/:id", handler.DeletePublisher)
}
