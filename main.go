package main

import (
	"log"

	"github.com/baijupadmanabhan/golang-fiber-api/database"
	"github.com/baijupadmanabhan/golang-fiber-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/api/Users", routes.CreateUser)
	app.Get("/api/Users", routes.GetAllUsers)
	app.Get("/api/Users/:id", routes.GetUserById)
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Fiber API")
}
