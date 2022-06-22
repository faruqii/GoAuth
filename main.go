package main

import (
	"github.com/faruqii/GoAuth/database"
	"github.com/faruqii/GoAuth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	})) // fronted can see what cookie we send

	routes.Setup(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
