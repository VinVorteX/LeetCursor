package handlers

import (
	"github.com/VinVorteX/LeetCursor/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetLeaderboard(c *fiber.Ctx) error {
	var users []models.User

	// Get query parameters for filtering and sorting
	platform := c.Query("platform", "all")
	sortBy := c.Query("sortBy", "totalSolved")
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 20)

	query := h.db.Model(&models.User{})

	// Apply filters
	if platform != "all" {
		query = query.Joins("JOIN coding_profiles ON users.id = coding_profiles.user_id").
			Where("coding_profiles.platform = ?", platform)
	}

	// Apply sorting
	switch sortBy {
	case "totalSolved":
		query = query.Order("total_solved DESC")
	case "rating":
		query = query.Order("rating DESC")
	}

	// Pagination
	offset := (page - 1) * limit

	var total int64
	query.Count(&total)

	err := query.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch leaderboard",
		})
	}

	return c.JSON(fiber.Map{
		"users": users,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}
