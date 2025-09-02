package routes

import (
	friendship "socio/controllers/friendships"

	"github.com/gofiber/fiber/v2"
)

func Friendships(r fiber.Router) {
	postsRoutes := r.Group("/friends")

	postsRoutes.Post("/", friendship.Add)
	postsRoutes.Get("/:id", friendship.Get)
	postsRoutes.Delete("/:id", friendship.Delete)
}
