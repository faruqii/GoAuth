package routes

import (
	"github.com/faruqii/GoAuth/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// User Auth
	app.Post("/api/register", controllers.Register)
	app.Post("/api/Login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	// Product
	app.Post("/api/CreateProduct", controllers.CreateProduct)
	app.Get("/api/GetAllProducts", controllers.GetAllProduct)
}
