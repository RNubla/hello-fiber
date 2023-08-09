package router

import (
	"github.com/RNubla/hello-fiber/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	v1 := api.Group("/v1", logger.New())
	v1.Get("/", handler.HelloV1)
}
