package middleware

import (
	"errors"
	"log"

	"github.com/faruqii/GoAuth/database"
	"github.com/faruqii/GoAuth/models"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Filter       func(c *fiber.Ctx) bool // Required
	Unauthorized fiber.Handler           // middleware specfic
}

func New(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.GetReqHeaders()

		if _, ok := header["Token"]; !ok {
			return errors.New("token missing")
		}

		// if exists check the user token table
		userToken := models.UserToken{}
		err := database.DB.Find(&userToken, "token = ?", header["Token"]).Error
		if err != nil {
			return errors.New("token fail")
		}
		c.Locals("user_id", userToken.UserID)
		log.Printf("user %d, pass token middleware", userToken.UserID)
		return c.Next()
	}
}
