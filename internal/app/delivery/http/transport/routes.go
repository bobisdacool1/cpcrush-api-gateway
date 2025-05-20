package transport

import (
	"github.com/gofiber/fiber/v3"

	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/delivery/http/handler"
)

func RegisterRoutes(app *fiber.App, hh *handler.HealthcheckHandler) {
	app.Get("/healthcheck", hh.Get)
}
