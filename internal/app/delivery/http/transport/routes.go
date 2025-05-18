package transport

import (
	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/delivery/http/handler"
)

func (app *Application) RegisterRoutes() {
	app.AppendEndpoint("/healthcheck", MethodGet, handler.GetHealthcheck)
}
