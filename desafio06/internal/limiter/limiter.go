package limiter

import (
	"context"
	"strings"
	"time"

	"github.com/gbuenodev/fullcycle_go_expert/desafio06/config"
	"github.com/gbuenodev/fullcycle_go_expert/desafio06/internal/storage"
)

type RateLimiter struct {
	storage storage.Storage
	config  *config.Config
}

func NewRateLimiter(storage storage.Storage, config *config.Config) *RateLimiter {
	return &RateLimiter{
		storage: storage,
		config:  config,
	}
}

func (rl *RateLimiter) Check(ctx context.Context, ip string, token string) (bool, error) {
	// Token takes priority over IP
	if token != "" {
		tier := parseTokenTier(token)
		limit := rl.getLimitForTier(tier)
		blockDuration := time.Duration(rl.config.RateLimitTokenBlockDuration) * time.Second
		key := "token:" + token

		return rl.checkLimit(ctx, key, limit, blockDuration)
	}

	limit := rl.config.RateLimitIP
	blockDuration := time.Duration(rl.config.RateLimitIPBlockDuration) * time.Second
	key := "ip:" + ip

	return rl.checkLimit(ctx, key, limit, blockDuration)
}

func (rl *RateLimiter) checkLimit(ctx context.Context, key string, limit int, blockDuration time.Duration) (bool, error) {
	blocked, err := rl.storage.IsBlocked(ctx, key)
	if err != nil {
		return false, err
	}

	if blocked {
		return false, nil
	}

	window := time.Duration(rl.config.RateLimitWindow) * time.Second
	count, err := rl.storage.Increment(ctx, key, window)
	if err != nil {
		return false, err
	}

	if count > int64(limit) {
		if err := rl.storage.Block(ctx, key, blockDuration); err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func parseTokenTier(token string) string {
	parts := strings.Split(token, ":")
	if len(parts) > 0 {
		tier := strings.ToLower(parts[0])
		if tier == "basic" || tier == "premium" {
			return tier
		}
	}
	return "basic"
}

func (rl *RateLimiter) getLimitForTier(tier string) int {
	switch tier {
	case "basic":
		return rl.config.RateLimitTokenBasic
	case "premium":
		return rl.config.RateLimitTokenPremium
	default:
		return rl.config.RateLimitTokenBasic
	}
}
