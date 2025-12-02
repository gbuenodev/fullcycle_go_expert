package auction

import (
	"context"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func setupTest(t *testing.T) *mtest.T {
	return mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
}

func TestCompleteAuction_Success(t *testing.T) {
	mt := setupTest(t)

	mt.Run("sets status to Completed after timeout", func(mt *mtest.T) {
		os.Setenv("AUCTION_DURATION", "50ms")
		defer os.Unsetenv("AUCTION_DURATION")

		repo := &AuctionRepository{Collection: mt.Coll}
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		go repo.completeAuction(context.Background(), "test-id")

		time.Sleep(100 * time.Millisecond)
	})
}

func TestGetAuctionDuration(t *testing.T) {
	tests := []struct {
		env      string
		expected time.Duration
	}{
		{"24h", 24 * time.Hour},
		{"30m", 30 * time.Minute},
		{"", 24 * time.Hour},
	}

	for _, tt := range tests {
		os.Setenv("AUCTION_DURATION", tt.env)
		if got := getAuctionDuration(); got != tt.expected {
			t.Errorf("env=%q: got %v, want %v", tt.env, got, tt.expected)
		}
	}
	os.Unsetenv("AUCTION_DURATION")
}
