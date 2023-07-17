package main

import (
	"github.com/alejlatorre/trivia-rest-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Post("/fact", handlers.CreateFact)
	app.Get("/list_facts", handlers.ListFacts)
}
