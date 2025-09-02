package server

import (
	"socio/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "socio/docs" // Swagger docs
)

// Default Error Handler
func errHandler(c *fiber.Ctx, e error) error {
	msg := e.Error()
	return c.Status(fiber.StatusInternalServerError).JSON(msg)
}

// Not found handler
var notFoundHandler = func(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON("Requested resource not found")
}

func addRoutes(app *fiber.App) {
	// Swagger UI route
	app.Get("/swagger/*", swagger.HandlerDefault)
	
	baseRouter := app.Group("/socio")

	routes.Users(baseRouter)
	routes.Friendships(baseRouter)
	routes.Posts(baseRouter)
}
