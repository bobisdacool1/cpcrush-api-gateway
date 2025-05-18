package handler_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"

	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/delivery/http/handler"
)

func TestGetHealthcheck_Unit(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})

	err := handler.GetHealthcheck(ctx)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, ctx.Response().StatusCode())
	assert.JSONEq(t, `{"status":"ok"}`, string(ctx.Response().Body()))
}

func TestGetHealthcheck(t *testing.T) {
	app := fiber.New()
	app.Get("/healthcheck", handler.GetHealthcheck)

	req := httptest.NewRequest("GET", "/healthcheck", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	expected := `{"status":"ok"}`
	buf := make([]byte, resp.ContentLength)
	_, _ = resp.Body.Read(buf)
	_ = resp.Body.Close()
	assert.JSONEq(t, expected, string(buf))
}
