// cmd/routes.go

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/bigfishfastswimer/rest-api-app/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Get("/api/v1/listfacts", handlers.ListFacts)
	app.Post("/api/v1/facts", handlers.CreateFact)
}
