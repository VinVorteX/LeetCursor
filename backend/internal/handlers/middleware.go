package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handler) AuthMiddleware(c *fiber.Ctx) error {
	// TODO: Implement JWT validation
	// For now, just pass through
	return c.Next()
} 