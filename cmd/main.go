package main

import (
	"go.uber.org/fx"

	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/config"
	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/delivery/http/transport"
	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/infrastructure/logger"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
			logger.NewLogger,
			transport.NewApplication,
		),
		fx.Invoke(
			func(app *transport.Application, lc fx.Lifecycle) {
				app.Start(lc)
			},
		),
	)

	app.Run()
}
