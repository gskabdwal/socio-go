package config

import (
	"socio/internals/database"
	friendships "socio/models/friendship"
	"socio/models/posts"
	"socio/models/users"
)

func Automigration() {
	database.Client().Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	database.Client().AutoMigrate(&users.Users{}, &friendships.Friendships{}, &posts.Posts{})
}
