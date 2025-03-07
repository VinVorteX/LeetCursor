package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/VinVorteX/LeetCursor/internal/config"
	"github.com/VinVorteX/LeetCursor/internal/database"
	"github.com/VinVorteX/LeetCursor/internal/handlers"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := database.Initialize(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize Redis
	redis, err := database.InitializeRedis(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}

	// Create Fiber app
	app := fiber.New()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// Initialize handlers
	h := handlers.NewHandler(db, redis)

	// Routes
	api := app.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/register", h.Register)
	auth.Post("/login", h.Login)
	auth.Get("/google", h.GoogleAuth)
	auth.Get("/google/callback", h.GoogleCallback)

	// Protected routes
	protected := api.Group("/", h.AuthMiddleware)
	protected.Get("/profile", h.GetProfile)
	protected.Put("/profile", h.UpdateProfile)
	protected.Get("/leaderboard", h.GetLeaderboard)

	// Start server
	log.Fatal(app.Listen(cfg.ServerPort))
}
