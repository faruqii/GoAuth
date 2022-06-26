package controllers

import (
	"github.com/faruqii/GoAuth/database"
	"github.com/faruqii/GoAuth/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CreateMerchant(c *fiber.Ctx) error {
	req := models.CreateMerchantParam{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	merchant := models.Merchant{
		Name: req.Name,
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
