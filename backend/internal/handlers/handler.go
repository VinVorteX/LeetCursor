package handlers

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Handler struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewHandler(db *gorm.DB, redis *redis.Client) *Handler {
	return &Handler{
		db:    db,
		redis: redis,
	}
} 