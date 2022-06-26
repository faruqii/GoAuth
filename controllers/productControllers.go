package controllers

import (
	"github.com/faruqii/GoAuth/database"
	"github.com/faruqii/GoAuth/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CreateProduct(c *fiber.Ctx) error {
	req := models.CreateProductParam{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	product := models.Product{
		Name:        req.Name,
		ProductType: req.ProductType,
		Price:       req.Price,
		MerchantID:  req.MerchantID,
	}
	merchant := models.Merchant{}
	// assign product to merchant and make sure merchant is not nil
	err := database.DB.Find(&merchant, "id = ?", req.MerchantID).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	err = database.DB.Create(&product).Error
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
	err := database.DB.Preload("Merchant").Find(&products).Error
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
