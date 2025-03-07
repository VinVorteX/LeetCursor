package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handler) GetProfile(c *fiber.Ctx) error {
	// TODO: Implement profile retrieval
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (h *Handler) UpdateProfile(c *fiber.Ctx) error {
	// TODO: Implement profile update
	return c.SendStatus(fiber.StatusNotImplemented)
} 