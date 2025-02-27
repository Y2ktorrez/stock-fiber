package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mutinho/src/dtos"
	"github.com/mutinho/src/services"
)

type ProductHandler struct {
	ProductService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: productService}
}

func (h *ProductHandler) RegisterRoutes(router fiber.Router) {
	router.Post("/product/create", h.CreateProduct)
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var req dtos.CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	productResponse, err := h.ProductService.CreateProduct(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"product": productResponse,
	})
}
