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
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetAllUsers)
	app.Get("/api/users/:id", routes.GetUserById)
	app.Put("/api/users/:id", routes.UpdateUser)
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Fiber API")
}
