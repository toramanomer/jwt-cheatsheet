package token

import (
	"sync"
)

type TokenStore struct {
	mu      sync.RWMutex
	revoked map[string]bool
}

func NewTokenStore() *TokenStore {
	return &TokenStore{
		revoked: make(map[string]bool),
	}
}

func (s *TokenStore) Revoke(jti string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.revoked[jti] = true
}

func (s *TokenStore) IsRevoked(jti string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.revoked[jti]
}
