package controllers

import (
	"net/http"

	"github.com/faruqii/GoAuth/database"
	"github.com/faruqii/GoAuth/models"
	"github.com/gofiber/fiber/v2"
)

func CreateMerchant(c *fiber.Ctx) error {
	// Create Merchant by models
	req := models.CreateMerchantParam{} // Request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	merchant := models.Merchant{
		Name: req.Name, // fill merchant with request body
	}
	err := database.DB.Create(&merchant).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Product created successfully",
		"data":    merchant,
	})
}

func GetAllMerchant(c *fiber.Ctx) error {
	// display all merchant
	// display all products that belong to that merchant
	var merchants []models.Merchant
	err := database.DB.Find(&merchants).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}
	products := models.Product{}
	for _, merchant := range merchants {
		err := database.DB.Where("merchant_id = ?", merchant.ID).Find(&products).Error
		if err != nil {
			return c.Status(http.StatusNotFound).JSON(err)
		}
		
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successful Get All Merchants",
		"data":    merchants,
	})

}

func SearchMerchantByName(c *fiber.Ctx) error {
	// In this function, we will search merchant and display all the products name that belong to that merchant
	// find merchant by name
	name := c.Params("name")
	merchant := []models.Merchant{}
	err := database.DB.Where("name = ?", name).Find(&merchant).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}
	// find all products that belong to that merchant then display it 
	products := []models.Product{}
	for _, merchant := range merchant {
		err := database.DB.Where("merchant_id = ?", merchant.ID).Find(&products).Error
		if err != nil {
			return c.Status(http.StatusNotFound).JSON(err)
		}
		merchant.Products = products
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successful Get All Merchants",
		"data":    merchant,
	})
	
}

func UpdateMerchant(c *fiber.Ctx) error {
	// Update Merchant by models
	req := models.CreateMerchantParam{} // Request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	merchant := models.Merchant{
		Name: req.Name, // fill merchant with request body
	}
	err := database.DB.Save(&merchant).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Product updated successfully",
		"data":    merchant,
	})
}

func DeleteMerchant(c *fiber.Ctx) error {
	// Delete Merchant by models
	id := c.Params("id")
	merchant := models.Merchant{}
	err := database.DB.Where("id = ?", id).Find(&merchant).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}
	err = database.DB.Delete(&merchant).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Product deleted successfully",
		"data":    merchant,
	})	
}
