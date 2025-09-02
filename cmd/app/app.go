package app

import (
	"log"

	"socio/internals/cache"
	"socio/internals/config"
	"socio/internals/database"
	"socio/internals/notifications"
	"socio/internals/server"
)

func Setup() {
	database.Connect()
	cache.Connect()
	config.Automigration()

	notifications.InitNotificationsSystem()
	notifications.Hydrate()

	server.Setup()
	app := server.New()

	if err := app.Listen(":3015"); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
