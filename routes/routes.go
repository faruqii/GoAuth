package routes

import (
	"github.com/faruqii/GoAuth/controllers"
	"github.com/faruqii/GoAuth/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// Endpoint
	api := app.Group("/api")

	// User Auth
	user := api.Group("/user")
	user.Post("/register", controllers.Register)
	user.Post("/Login", controllers.Login)
	user.Get("/user", controllers.User)
	user.Post("/logout", controllers.Logout)

	// Product
	product := api.Group("/product")
	product.Use(middleware.New(middleware.Config{})).Post("/CreateProduct", controllers.CreateProduct)
	product.Get("/GetAllProducts", controllers.GetAllProduct)

	// Merchant
	merchant := api.Group("/merchant")
	merchant.Post("/CreateMerchant", controllers.CreateMerchant)
	merchant.Get("/GetAllMerchant", controllers.GetAllMerchant)
}
