package controllers

import (
	"github.com/faruqii/GoAuth/database"
	"github.com/faruqii/GoAuth/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CreateProduct(c *fiber.Ctx) error {

	// v := validator.New()
	product := models.Product{}
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	// err := v.Struct(product)

	// for _, e := range err.(validator.ValidationErrors) {
	// 	fmt.Println(e)
	// }

	err := database.DB.Create(&product).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	response := fiber.Map{
		"status":  "success",
		"message": "Product created successfully",
		"data":    product,
	}

	return c.Status(http.StatusCreated).JSON(response)
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
