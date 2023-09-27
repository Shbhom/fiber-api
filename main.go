package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shbhom/fiber-api/database"
	"github.com/shbhom/fiber-api/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("welcome users to my first go project")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", welcome)
	app.Get("/users", routes.GetUsers)
	app.Get("/user/:id", routes.GetUserById)
	app.Post("/user", routes.CreateUser)
	app.Put("/user/:id", routes.UpdateUser)
}

func main() {
	app := fiber.New()
	database.ConnectDB()

	setupRoutes(app)

	log.Fatal(app.Listen(":5500"))
}
