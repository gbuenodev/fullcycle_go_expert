package auction

import (
	"context"
	"os"
	"testing"
	"time"

	auctionentity "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/entity/auction_entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func setupTest(t *testing.T) *mtest.T {
	return mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
}

func TestCompleteAuction_Success(t *testing.T) {
	mt := setupTest(t)

	mt.Run("tests complete auction", func(mt *mtest.T) {
		os.Setenv("AUCTION_DURATION", "50ms")
		defer os.Unsetenv("AUCTION_DURATION")

		auction, err := auctionentity.NewAuction(
			"test-product",
			"test-category",
			"test-description",
			auctionentity.New,
		)
		if err != nil {
			t.Fatalf("NewAuction() error = %v", err)
		}

		mt.AddMockResponses(
			mtest.CreateSuccessResponse(),
			mtest.CreateSuccessResponse(),
			mtest.CreateCursorResponse(1, "auctions.auctions", mtest.FirstBatch, bson.D{
				{Key: "_id", Value: auction.ID},
				{Key: "product_name", Value: auction.ProductName},
				{Key: "category", Value: auction.Category},
				{Key: "description", Value: auction.Description},
				{Key: "condition", Value: auction.Condition},
				{Key: "status", Value: auctionentity.Completed},
				{Key: "timestamp", Value: auction.Timestamp.Unix()},
			}),
		)

		repo := &AuctionRepository{Collection: mt.Coll}

		err = repo.CreateAuction(context.Background(), auction)
		if err != nil {
			t.Fatalf("CreateAuction() error = %v", err)
		}

		time.Sleep(100 * time.Millisecond)

		completedAuction, err := repo.FindAuctionById(context.Background(), auction.ID)
		if err != nil {
			t.Fatalf("FindAuctionById() error = %v", err)
		}

		if completedAuction.Status != auctionentity.Completed {
			t.Errorf("Auction status = %v, want %v", completedAuction.Status, auctionentity.Completed)
		}
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
