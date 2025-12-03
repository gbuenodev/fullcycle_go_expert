package limiter

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/gbuenodev/fullcycle_go_expert/desafio06/config"
	"github.com/gbuenodev/fullcycle_go_expert/desafio06/internal/storage"
)

func TestParseTokenTier(t *testing.T) {
	tests := []struct {
		name     string
		token    string
		expected string
	}{
		{
			name:     "valid basic token",
			token:    "basic:550e8400-e29b-41d4-a716-446655440000",
			expected: "basic",
		},
		{
			name:     "valid premium token",
			token:    "premium:550e8400-e29b-41d4-a716-446655440000",
			expected: "premium",
		},
		{
			name:     "uppercase basic token",
			token:    "BASIC:550e8400-e29b-41d4-a716-446655440000",
			expected: "basic",
		},
		{
			name:     "uppercase premium token",
			token:    "PREMIUM:550e8400-e29b-41d4-a716-446655440000",
			expected: "premium",
		},
		{
			name:     "invalid tier defaults to basic",
			token:    "invalid:550e8400-e29b-41d4-a716-446655440000",
			expected: "basic",
		},
		{
			name:     "empty token defaults to basic",
			token:    "",
			expected: "basic",
		},
		{
			name:     "malformed token no colon",
			token:    "justAToken",
			expected: "basic",
		},
		{
			name:     "token with multiple colons uses first part",
			token:    "premium:part2:part3",
			expected: "premium",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseTokenTier(tt.token)
			if result != tt.expected {
				t.Errorf("parseTokenTier(%q) = %q; want %q", tt.token, result, tt.expected)
			}
		})
	}
}

func TestGetLimitForTier(t *testing.T) {
	cfg := &config.Config{
		RateLimitTokenBasic:   50,
		RateLimitTokenPremium: 200,
	}

	rl := NewRateLimiter(nil, cfg)

	tests := []struct {
		name     string
		tier     string
		expected int
	}{
		{
			name:     "basic tier",
			tier:     "basic",
			expected: 50,
		},
		{
			name:     "premium tier",
			tier:     "premium",
			expected: 200,
		},
		{
			name:     "unknown tier defaults to basic",
			tier:     "unknown",
			expected: 50,
		},
		{
			name:     "empty tier defaults to basic",
			tier:     "",
			expected: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rl.getLimitForTier(tt.tier)
			if result != tt.expected {
				t.Errorf("getLimitForTier(%q) = %d; want %d", tt.tier, result, tt.expected)
			}
		})
	}
}

func TestRateLimiter_Check_TokenPriority(t *testing.T) {
	mockStorage := storage.NewMockStorage()
	cfg := &config.Config{
		RateLimitIP:                 10,
		RateLimitIPBlockDuration:    300,
		RateLimitTokenBasic:         50,
		RateLimitTokenPremium:       200,
		RateLimitTokenBlockDuration: 300,
		RateLimitWindow:             1,
	}

	rl := NewRateLimiter(mockStorage, cfg)
	ctx := context.Background()

	tests := []struct {
		name          string
		ip            string
		token         string
		expectedKey   string
		expectedLimit int
	}{
		{
			name:          "token present uses token limit",
			ip:            "192.168.1.1",
			token:         "premium:abc-123",
			expectedKey:   "token:premium:abc-123",
			expectedLimit: 200,
		},
		{
			name:          "no token uses IP limit",
			ip:            "192.168.1.1",
			token:         "",
			expectedKey:   "ip:192.168.1.1",
			expectedLimit: 10,
		},
		{
			name:          "basic token uses basic limit",
			ip:            "192.168.1.1",
			token:         "basic:xyz-789",
			expectedKey:   "token:basic:xyz-789",
			expectedLimit: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage.Reset()

			allowed, err := rl.Check(ctx, tt.ip, tt.token)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !allowed {
				t.Error("expected request to be allowed")
			}

			if len(mockStorage.IncrementCalls) != 1 {
				t.Fatalf("expected 1 Increment call, got %d", len(mockStorage.IncrementCalls))
			}

			if mockStorage.IncrementCalls[0].Key != tt.expectedKey {
				t.Errorf("Increment called with key %q; want %q",
					mockStorage.IncrementCalls[0].Key, tt.expectedKey)
			}
		})
	}
}

func TestRateLimiter_CheckLimit_Logic(t *testing.T) {
	tests := []struct {
		name            string
		setupMock       func(*storage.MockStorage)
		limit           int
		expectAllowed   bool
		expectBlockCall bool
		expectError     bool
	}{
		{
			name: "request under limit is allowed",
			setupMock: func(m *storage.MockStorage) {
				m.SetCounter("test-key", 5)
			},
			limit:           10,
			expectAllowed:   true,
			expectBlockCall: false,
		},
		{
			name: "request at limit is allowed",
			setupMock: func(m *storage.MockStorage) {
				m.SetCounter("test-key", 9)
			},
			limit:           10,
			expectAllowed:   true,
			expectBlockCall: false,
		},
		{
			name: "request over limit is blocked",
			setupMock: func(m *storage.MockStorage) {
				m.SetCounter("test-key", 10)
			},
			limit:           10,
			expectAllowed:   false,
			expectBlockCall: true,
		},
		{
			name: "already blocked entity is denied",
			setupMock: func(m *storage.MockStorage) {
				m.SetBlocked("test-key", true)
			},
			limit:           10,
			expectAllowed:   false,
			expectBlockCall: false,
		},
		{
			name: "IsBlocked error is propagated",
			setupMock: func(m *storage.MockStorage) {
				m.IsBlockedError = errors.New("redis connection failed")
			},
			limit:       10,
			expectError: true,
		},
		{
			name: "Increment error is propagated",
			setupMock: func(m *storage.MockStorage) {
				m.IncrementError = errors.New("redis increment failed")
			},
			limit:       10,
			expectError: true,
		},
		{
			name: "Block error is propagated",
			setupMock: func(m *storage.MockStorage) {
				m.SetCounter("test-key", 10)
				m.BlockError = errors.New("redis block failed")
			},
			limit:       10,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := storage.NewMockStorage()
			tt.setupMock(mockStorage)

			cfg := &config.Config{
				RateLimitIP:              tt.limit,
				RateLimitIPBlockDuration: 300,
				RateLimitWindow:          1,
			}

			rl := NewRateLimiter(mockStorage, cfg)
			ctx := context.Background()

			allowed, err := rl.checkLimit(ctx, "test-key", tt.limit, 300*time.Second)

			if tt.expectError {
				if err == nil {
					t.Error("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if allowed != tt.expectAllowed {
				t.Errorf("expected allowed=%v, got %v", tt.expectAllowed, allowed)
			}

			if tt.expectBlockCall {
				if len(mockStorage.BlockCalls) == 0 {
					t.Error("expected Block to be called but it wasn't")
				}
			} else if tt.name != "already blocked entity is denied" {
				// Only check if not the "already blocked" test case
				if len(mockStorage.BlockCalls) > 0 {
					t.Error("expected Block not to be called but it was")
				}
			}
		})
	}
}

func TestRateLimiter_Check_Integration(t *testing.T) {
	mockStorage := storage.NewMockStorage()
	cfg := &config.Config{
		RateLimitIP:                 3,
		RateLimitIPBlockDuration:    300,
		RateLimitTokenBasic:         5,
		RateLimitTokenPremium:       10,
		RateLimitTokenBlockDuration: 300,
		RateLimitWindow:             1,
	}

	rl := NewRateLimiter(mockStorage, cfg)
	ctx := context.Background()

	t.Run("multiple requests respect IP limit", func(t *testing.T) {
		mockStorage.Reset()
		ip := "192.168.1.100"

		// First 3 requests should pass
		for i := 1; i <= 3; i++ {
			allowed, err := rl.Check(ctx, ip, "")
			if err != nil {
				t.Fatalf("request %d: unexpected error: %v", i, err)
			}
			if !allowed {
				t.Errorf("request %d: expected to be allowed", i)
			}
		}

		// 4th request should be blocked
		allowed, err := rl.Check(ctx, ip, "")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if allowed {
			t.Error("request 4: expected to be blocked")
		}

		// Subsequent requests should also be blocked
		allowed, err = rl.Check(ctx, ip, "")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if allowed {
			t.Error("request 5: expected to be blocked (entity is blocked)")
		}
	})

	t.Run("multiple requests respect token limit", func(t *testing.T) {
		mockStorage.Reset()
		token := "basic:test-token-123"

		// First 5 requests should pass (basic limit)
		for i := 1; i <= 5; i++ {
			allowed, err := rl.Check(ctx, "192.168.1.200", token)
			if err != nil {
				t.Fatalf("request %d: unexpected error: %v", i, err)
			}
			if !allowed {
				t.Errorf("request %d: expected to be allowed", i)
			}
		}

		// 6th request should be blocked
		allowed, err := rl.Check(ctx, "192.168.1.200", token)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if allowed {
			t.Error("request 6: expected to be blocked")
		}
	})

	t.Run("different keys are tracked independently", func(t *testing.T) {
		mockStorage.Reset()

		ip1 := "192.168.1.1"
		ip2 := "192.168.1.2"

		// Max out IP1
		for i := 1; i <= 3; i++ {
			rl.Check(ctx, ip1, "")
		}

		// IP1 should be blocked
		allowed, err := rl.Check(ctx, ip1, "")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if allowed {
			t.Error("IP1 should be blocked")
		}

		// IP2 should still work
		allowed, err = rl.Check(ctx, ip2, "")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !allowed {
			t.Error("IP2 should be allowed (different key)")
		}
	})

	t.Run("window duration is passed correctly", func(t *testing.T) {
		mockStorage.Reset()

		rl.Check(ctx, "192.168.1.1", "")

		if len(mockStorage.IncrementCalls) != 1 {
			t.Fatalf("expected 1 Increment call, got %d", len(mockStorage.IncrementCalls))
		}

		expectedWindow := time.Duration(cfg.RateLimitWindow) * time.Second
		if mockStorage.IncrementCalls[0].Expiration != expectedWindow {
			t.Errorf("expected window duration %v, got %v",
				expectedWindow, mockStorage.IncrementCalls[0].Expiration)
		}
	})

	t.Run("block duration is passed correctly for IP", func(t *testing.T) {
		mockStorage.Reset()
		mockStorage.SetCounter("ip:192.168.1.1", 3)

		rl.Check(ctx, "192.168.1.1", "")

		if len(mockStorage.BlockCalls) != 1 {
			t.Fatalf("expected 1 Block call, got %d", len(mockStorage.BlockCalls))
		}

		expectedDuration := time.Duration(cfg.RateLimitIPBlockDuration) * time.Second
		if mockStorage.BlockCalls[0].Expiration != expectedDuration {
			t.Errorf("expected block duration %v, got %v",
				expectedDuration, mockStorage.BlockCalls[0].Expiration)
		}
	})

	t.Run("block duration is passed correctly for token", func(t *testing.T) {
		mockStorage.Reset()
		token := "premium:test-123"
		mockStorage.SetCounter("token:"+token, 10)

		rl.Check(ctx, "192.168.1.1", token)

		if len(mockStorage.BlockCalls) != 1 {
			t.Fatalf("expected 1 Block call, got %d", len(mockStorage.BlockCalls))
		}

		expectedDuration := time.Duration(cfg.RateLimitTokenBlockDuration) * time.Second
		if mockStorage.BlockCalls[0].Expiration != expectedDuration {
			t.Errorf("expected block duration %v, got %v",
				expectedDuration, mockStorage.BlockCalls[0].Expiration)
		}
	})
}
