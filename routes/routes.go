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
	product := api.Group("/product").Use(middleware.New(middleware.Config{}))
	product.Post("/CreateProduct", controllers.CreateProduct)
	product.Get("/GetAllProducts", controllers.GetAllProduct)
	product.Get("/SearchProductByName", controllers.SearchProductByName)
	product.Patch("/UpdateProduct", controllers.UpdateProduct)
	product.Delete("/DeleteProduct", controllers.DeleteProduct)

	// Merchant
	merchant := api.Group("/merchant").Use(middleware.New(middleware.Config{}))
	merchant.Post("/CreateMerchant", controllers.CreateMerchant)
	merchant.Get("/GetAllMerchant", controllers.GetAllMerchant)
	merchant.Get("/SearchMerchantByName", controllers.SearchMerchantByName)
	merchant.Patch("/UpdateMerchant", controllers.UpdateMerchant)
	merchant.Delete("/DeleteMerchant", controllers.DeleteMerchant)
}
