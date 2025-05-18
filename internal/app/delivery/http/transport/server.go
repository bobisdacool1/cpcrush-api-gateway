package transport

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"

	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/config"
)

type (
	Application struct {
		Name    string
		Version string

		Config config.AppConfig
		Server *fiber.App
	}

	HttpMethod string
)

const (
	MethodGet    HttpMethod = http.MethodGet
	MethodPost   HttpMethod = http.MethodPost
	MethodPut    HttpMethod = http.MethodPut
	MethodPatch  HttpMethod = http.MethodPatch
	MethodDelete HttpMethod = http.MethodDelete
)

func NewApplication(lc fx.Lifecycle, cfg config.AppConfig) *Application {
	app := &Application{
		Server: fiber.New(),
		Config: cfg,
	}

	app.RegisterRoutes()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Server.Listen(":1488")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Server.Shutdown()
		},
	})

	return app
}

func (app *Application) Start() error {
	return app.Server.Listen(":1488")
}

func (app *Application) AppendEndpoint(path string, method HttpMethod, handler fiber.Handler) {
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
