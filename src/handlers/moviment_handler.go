package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mutinho/src/dtos"
	"github.com/mutinho/src/services"
)

type MovimentHandler struct {
	movimentService *services.MovimentService
}

func NewMovimentHandler(movimentService *services.MovimentService) *MovimentHandler {
	return &MovimentHandler{movimentService: movimentService}
}

func (h *MovimentHandler) RegisterRoutes(router fiber.Router) {
	router.Post("/moviment", h.CreateMoviment)
	router.Get("/moviments", h.GetAllMoviments)
}

func (h *MovimentHandler) CreateMoviment(c *fiber.Ctx) error {
	var req dtos.CreateMovimentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	movimentResponse, err := h.movimentService.CreateMoviment(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create moviment",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Moviment created successfully",
		"moviment": movimentResponse,
	})
}

func (h *MovimentHandler) GetAllMoviments(c *fiber.Ctx) error {
	movimentResponses, err := h.movimentService.GetAllMoviments()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve moviments",
		})
	}

	return c.Status(fiber.StatusOK).JSON(movimentResponses)
}
