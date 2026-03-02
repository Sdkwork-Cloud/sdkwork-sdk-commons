package common

import (
	"time"
	"sync"
)

type AuthMode string

const (
	AuthModeNone   AuthMode = "none"
	AuthModeAPIKey AuthMode = "api_key"
	AuthModeBearer AuthMode = "bearer"
	AuthModeOAuth2 AuthMode = "oauth2"
)

type AuthTokens struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    *int64
	TokenType    string
	Scope        string
	Extra        map[string]interface{}
}

type AuthTokenManager struct {
	mu             sync.Mutex
	tokens         *AuthTokens
	refreshCallback func() (*AuthTokens, error)
}

func NewAuthTokenManager() *AuthTokenManager {
	return &AuthTokenManager{}
}

func (m *AuthTokenManager) SetTokens(tokens *AuthTokens) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tokens = tokens
}

func (m *AuthTokenManager) GetTokens() *AuthTokens {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.tokens
}

func (m *AuthTokenManager) GetAccessToken() string {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.tokens != nil {
		return m.tokens.AccessToken
	}
	return ""
}

func (m *AuthTokenManager) IsTokenValid() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.tokens == nil {
		return false
	}
	if m.tokens.ExpiresAt == nil {
		return true
	}
	return time.Now().UnixMilli() < *m.tokens.ExpiresAt
}

func (m *AuthTokenManager) RequiresRefresh() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.tokens == nil || m.tokens.RefreshToken == "" || m.tokens.ExpiresAt == nil {
		return false
	}
	buffer := int64(5 * 60 * 1000)
	return time.Now().UnixMilli() >= (*m.tokens.ExpiresAt - buffer)
}

func (m *AuthTokenManager) SetRefreshCallback(callback func() (*AuthTokens, error)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.refreshCallback = callback
}

func (m *AuthTokenManager) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tokens = nil
}

func BuildAuthHeaders(token string, mode AuthMode) map[string]string {
	if mode == AuthModeBearer {
		return map[string]string{"Authorization": "Bearer " + token}
	} else if mode == AuthModeAPIKey {
		return map[string]string{"Authorization": "Bearer " + token}
	}
	return map[string]string{}
}

func IsTokenValid(tokens *AuthTokens) bool {
	if tokens == nil {
		return false
	}
	if tokens.ExpiresAt == nil {
		return true
	}
	return time.Now().UnixMilli() < *tokens.ExpiresAt
}

func RequiresRefresh(tokens *AuthTokens) bool {
	if tokens == nil || tokens.RefreshToken == "" || tokens.ExpiresAt == nil {
		return false
	}
	buffer := int64(5 * 60 * 1000)
	return time.Now().UnixMilli() >= (*tokens.ExpiresAt - buffer)
}
