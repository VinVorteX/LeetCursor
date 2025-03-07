package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handler) Register(c *fiber.Ctx) error {
	// TODO: Implement registration
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	// TODO: Implement login
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (h *Handler) GoogleAuth(c *fiber.Ctx) error {
	// TODO: Implement Google OAuth
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (h *Handler) GoogleCallback(c *fiber.Ctx) error {
	// TODO: Implement Google OAuth callback
	return c.SendStatus(fiber.StatusNotImplemented)
} 