package routes

import (
	"labs/apis"

	"github.com/gofiber/fiber/v2"
)

func AuthorRoute(app *fiber.App) {
	app.Post("/Author", apis.CreateAuthor)
	app.Get("/Author/:ID", apis.GetAAuthor)
}
