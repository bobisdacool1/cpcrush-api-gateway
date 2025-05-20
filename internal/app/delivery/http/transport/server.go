package transport

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/config"
)

type (
	Application struct {
		Config config.AppConfig
		Server *fiber.App
		Logger *zap.Logger
	}
)

func NewApplication(cfg config.AppConfig, logger *zap.Logger, server *fiber.App) *Application {
	return &Application{
		Server: server,
		Config: cfg,
		Logger: logger,
	}
}

func (app *Application) Start(lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := app.Server.Listen(app.Config.Port); err != nil {
					app.Logger.Fatal("Fiber server failed", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Server.Shutdown()
		},
	})
}
