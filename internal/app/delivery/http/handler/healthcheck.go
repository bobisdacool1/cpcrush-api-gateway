package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func GetHealthcheck(c fiber.Ctx) error {
	c.Status(http.StatusOK)

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
