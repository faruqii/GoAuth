package controllers

import (
	"github.com/faruqii/GoAuth/database"
	"github.com/faruqii/GoAuth/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CreateProduct(c *fiber.Ctx) error {

	product := models.Product{}
	merchant := models.Merchant{}
	if err := c.BodyParser(&product); err != nil {
		// assign product to merchant and make sure merchant is not nil
		database.DB.Find(&merchant, "name = ?", product.Merchant)
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	err := database.DB.Create(&product).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Product created successfully",
		"data":    product,
	})
}

func GetAllProduct(c *fiber.Ctx) error {

	var products []models.Product
	err := database.DB.Find(&products).Error
	if err != nil {
		return c.JSON(err)
	}
	response := fiber.Map{
		"status":  "success",
		"message": "Successful Get All Products",
		"data":    products,
	}
	return c.Status(http.StatusOK).JSON(response)
}
