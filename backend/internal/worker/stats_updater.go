package worker

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/VinVorteX/LeetCursor/internal/models"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

const (
	TypeUpdateUserStats = "stats:update"
)

type LeetCodeStats struct {
	Easy   int
	Medium int
	Hard   int
	Total  int
}

func fetchLeetCodeStats(username string) (*LeetCodeStats, error) {
	// TODO: Implement actual LeetCode API call
	return &LeetCodeStats{
		Easy:   10,
		Medium: 5,
		Hard:   2,
		Total:  17,
	}, nil
}

type StatsUpdater struct {
	client *asynq.Client
	db     *gorm.DB
}

func NewStatsUpdater(redisAddr string, db *gorm.DB) *StatsUpdater {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	return &StatsUpdater{client: client, db: db}
}

func (s *StatsUpdater) Start(redisAddr string) {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeUpdateUserStats, s.handleUpdateUserStats)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}

func (s *StatsUpdater) handleUpdateUserStats(ctx context.Context, t *asynq.Task) error {
	var user models.User
	if err := json.Unmarshal(t.Payload(), &user); err != nil {
		return err
	}

	// Update LeetCode stats
	if user.LeetcodeID != "" {
		stats, err := fetchLeetCodeStats(user.LeetcodeID)
		if err == nil {
			user.EasySolved = stats.Easy
			user.MediumSolved = stats.Medium
			user.HardSolved = stats.Hard
			user.TotalSolved = stats.Total
		}
	}

	// Update other platforms...

	user.LastUpdated = time.Now()
	return s.db.Save(&user).Error
}
