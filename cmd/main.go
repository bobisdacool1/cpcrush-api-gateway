package main

import (
	"go.uber.org/fx"

	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/config"
	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/delivery/http/transport"
)

func main() {
	app := fx.New(
		fx.Provide(transport.NewApplication),
		fx.Provide(config.NewConfig),
		fx.Invoke(func(app *transport.Application) {
			app.Start()
		}),
	)

	app.Run()
}
