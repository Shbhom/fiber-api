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
	app.Post("/user", routes.CreateUser)
	app.Get("/users", routes.GetUsers)
	app.Get("/user/:id", routes.GetUserById)
	app.Put("/user/:id", routes.UpdateUser)
	app.Delete("/user/:id", routes.DeleteUser)

	app.Post("/product", routes.CreateProduct)
	app.Get("/products", routes.GetProducts)
	app.Get("/product/:id", routes.GetProductById)
	app.Put("/product/:id", routes.UpdateProduct)
	app.Delete("/product/:id", routes.DeleteProduct)

	app.Post("/order", routes.CreateOrder)
	app.Get("/order/:id", routes.GetOrderById)
}

func main() {
	app := fiber.New()
	database.ConnectDB()

	setupRoutes(app)

	log.Fatal(app.Listen(":5500"))
}
