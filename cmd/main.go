// Package main Social Media Backend API
//
// This is a social media backend API server.
//
//	Title: Social Media Backend API
//	Description: A REST API for social media functionality including users, posts, and friendships
//	Version: 1.0.0
//	Host: localhost:3015
//	BasePath: /socio
//	Schemes: http, https
//
//	Contact:
//	Name: API Support
//	Email: support@socialmedia.com
//
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
package main

import (
	"fmt"

	"socio/cmd/app"
)

func main() {
	fmt.Println("Application is booting...")
	app.Setup()
}
