package controllers

import (
	"net/http"

	"github.com/faruqii/GoAuth/database"
	"github.com/faruqii/GoAuth/models"
	"github.com/gofiber/fiber/v2"
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
		return c.Status(http.StatusNotFound).JSON(err.Error())
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

	products := []models.Product{}
	err := database.DB.Preload("Merchant").Find(&products).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successful Get All Products",
		"data":    products,
	})
}

func SearchProductByName(c *fiber.Ctx) error {
	// find product by name using query Like from postgres
	name := c.Params("name")
	products := models.Product{}
	err := database.DB.Where("name LIKE ?", "%"+name+"%").Find(&products).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successful Search Product By Name",
		"data":    products,
	})

}

func UpdateProduct(c *fiber.Ctx) error {
	// Update Product
	id := c.Params("id")
	product := models.Product{}
	err := database.DB.Find(&product, "id = ?", id).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}
	req := models.CreateProductParam{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	product.Name = req.Name
	product.ProductType = req.ProductType
	product.Price = req.Price
	err = database.DB.Save(&product).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successful Update Product",
		"data":    product,
	})

}

func DeleteProduct(c *fiber.Ctx) error {
	// Delete Product
	id := c.Params("id")
	product := models.Product{}
	err := database.DB.Find(&product, "id = ?", id).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}
	err = database.DB.Delete(&product).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successful Delete Product",
		"data":    product,
	})
}
