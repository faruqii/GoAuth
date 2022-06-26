package controllers

import (
	"github.com/faruqii/GoAuth/database"
	"github.com/faruqii/GoAuth/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CreateMerchant(c *fiber.Ctx) error {
	merchant := models.Merchant{}
	if err := c.BodyParser(&merchant); err != nil {
		// assign product to merchant and make sure merchant is not nil
		return c.Status(http.StatusBadRequest).JSON(err.Error())
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

	var merchant []models.Merchant
	err := database.DB.Find(&merchant).Error
	if err != nil {
		return c.JSON(err)
	}
	response := fiber.Map{
		"status":  "success",
		"message": "Successful Get All Merchant",
		"data":    merchant,
	}
	return c.Status(http.StatusOK).JSON(response)
}
