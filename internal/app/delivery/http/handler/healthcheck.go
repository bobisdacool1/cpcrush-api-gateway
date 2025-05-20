package handler

import (
	"github.com/gofiber/fiber/v3"

	"github.com/bobisdacool1/cpcrush/api-gateway/internal/app/usecase/healthcheck"
)

type HealthcheckHandler struct {
	service healthcheck.Service
}

func NewHealthcheckHandler(s healthcheck.Service) *HealthcheckHandler {
	return &HealthcheckHandler{service: s}
}

func (h *HealthcheckHandler) Get(c fiber.Ctx) error {
	return c.JSON(h.service.Ping())
}
