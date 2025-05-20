package testutils

import (
	"context"
	"testing"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"

	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/config"
	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/delivery/http/handler"
	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/delivery/http/transport"
	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/infrastructure/logger"
	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/usecase/healthcheck"
)

func NewTestApplication(t *testing.T) *transport.Application {
	t.Helper()

	var app *transport.Application
	var hh *handler.HealthcheckHandler

	fxApp := fx.New(
		fx.Provide(
			config.MustNewConfig,
			logger.NewLogger,
			healthcheck.NewService,
			handler.NewHealthcheckHandler,
			transport.NewApplication,
			fiber.New,
		),
		fx.Populate(&app, &hh),
		fx.Invoke(
			transport.RegisterRoutes,
			func(app *transport.Application, lc fx.Lifecycle) {
				app.Start(lc)
			},
		),
	)

	if err := fxApp.Start(context.Background()); err != nil {
		t.Fatalf("failed to start fxApp: %v", err)
	}

	t.Cleanup(func() {
		if err := fxApp.Stop(context.Background()); err != nil {
			t.Logf("failed to stop fxApp: %v", err)
		}
	})

	return app
}
