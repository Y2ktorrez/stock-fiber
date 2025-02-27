package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mutinho/src/dtos"
	"github.com/mutinho/src/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterRoutes(router fiber.Router) {
	router.Post("/user/register", h.CreateUser)
	router.Get("/user/email/:email", h.GetUserByEmail)
	router.Get("/user/slug/:slug", h.GetUserBySlug)
	router.Get("/users", h.GetAllUsers)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req dtos.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	userResponse, err := h.userService.CreateUser(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    userResponse,
	})
}

func (h *UserHandler) GetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email is required",
		})
	}

	userResponse, err := h.userService.GetUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(userResponse)
}

func (h *UserHandler) GetUserBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Slug is required",
		})
	}

	userResponse, err := h.userService.GetUserBySlug(slug)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(userResponse)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	userResponses, err := h.userService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}

	return c.Status(fiber.StatusOK).JSON(userResponses)
}
