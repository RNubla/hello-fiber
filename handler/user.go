package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{"status": http.StatusOK, "user": id})
}
