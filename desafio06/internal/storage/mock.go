package storage

import (
	"context"
	"sync"
	"time"
)

type MockStorage struct {
	mu sync.RWMutex

	// State
	counters map[string]int64
	blocked  map[string]bool

	// Control behavior
	IncrementError error
	IsBlockedError error
	BlockError     error

	// Call tracking
	IncrementCalls []IncrementCall
	IsBlockedCalls []IsBlockedCall
	BlockCalls     []BlockCall
}

type IncrementCall struct {
	Key        string
	Expiration time.Duration
}

type IsBlockedCall struct {
	Key string
}

type BlockCall struct {
	Key        string
	Expiration time.Duration
}

func NewMockStorage() *MockStorage {
	return &MockStorage{
		counters: make(map[string]int64),
		blocked:  make(map[string]bool),
	}
}

func (m *MockStorage) Increment(ctx context.Context, key string, expiration time.Duration) (int64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.IncrementCalls = append(m.IncrementCalls, IncrementCall{
		Key:        key,
		Expiration: expiration,
	})

	if m.IncrementError != nil {
		return 0, m.IncrementError
	}

	m.counters[key]++
	return m.counters[key], nil
}

func (m *MockStorage) IsBlocked(ctx context.Context, key string) (bool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	m.IsBlockedCalls = append(m.IsBlockedCalls, IsBlockedCall{
		Key: key,
	})

	if m.IsBlockedError != nil {
		return false, m.IsBlockedError
	}

	return m.blocked[key], nil
}

func (m *MockStorage) Block(ctx context.Context, key string, expiration time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.BlockCalls = append(m.BlockCalls, BlockCall{
		Key:        key,
		Expiration: expiration,
	})

	if m.BlockError != nil {
		return m.BlockError
	}

	m.blocked[key] = true
	return nil
}

// Helper methods for testing
func (m *MockStorage) SetBlocked(key string, blocked bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.blocked[key] = blocked
}

func (m *MockStorage) SetCounter(key string, count int64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.counters[key] = count
}

func (m *MockStorage) Reset() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.counters = make(map[string]int64)
	m.blocked = make(map[string]bool)
	m.IncrementCalls = nil
	m.IsBlockedCalls = nil
	m.BlockCalls = nil
	m.IncrementError = nil
	m.IsBlockedError = nil
	m.BlockError = nil
}
