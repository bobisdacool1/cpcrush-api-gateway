package transport

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/config"
)

type (
	Application struct {
		Name    string
		Version string

		Config config.AppConfig
		Server *fiber.App
		Logger *zap.Logger
	}

	HTTPMethod string
)

const (
	MethodGet    HTTPMethod = http.MethodGet
	MethodPost   HTTPMethod = http.MethodPost
	MethodPut    HTTPMethod = http.MethodPut
	MethodPatch  HTTPMethod = http.MethodPatch
	MethodDelete HTTPMethod = http.MethodDelete
)

func NewApplication(cfg config.AppConfig, logger *zap.Logger) *Application {
	return &Application{
		Server: fiber.New(),
		Config: cfg,
		Logger: logger,
	}
}

func (app *Application) Start(lc fx.Lifecycle) {
	app.RegisterRoutes()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				addr := ":1488"
				app.Logger.Info("Starting Fiber server", zap.String("addr", addr))
				if err := app.Server.Listen(addr); err != nil {
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

func (app *Application) AppendEndpoint(path string, method HTTPMethod, handler fiber.Handler) {
	switch method {
	case MethodGet:
		app.Server.Get(path, handler)
	case MethodPost:
		app.Server.Post(path, handler)
	case MethodPut:
		app.Server.Put(path, handler)
	case MethodPatch:
		app.Server.Patch(path, handler)
	case MethodDelete:
		app.Server.Delete(path, handler)
	}
}
