package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "code": http.StatusOK})
}

func HelloV1(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "code": http.StatusOK, "route": c.Route().Path})
}
