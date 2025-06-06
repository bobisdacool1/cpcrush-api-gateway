package main

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"

	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/config"
	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/delivery/http/handler"
	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/delivery/http/transport"
	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/infrastructure/logger"
	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/usecase/healthcheck"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.MustNewConfig,
			logger.NewLogger,
			healthcheck.NewService,
			handler.NewHealthcheckHandler,
			transport.NewApplication,
			fiber.New,
		),
		fx.Invoke(
			transport.RegisterRoutes,
			func(app *transport.Application, lc fx.Lifecycle) {
				app.Start(lc)
			},
		),
	)

	app.Run()
}
