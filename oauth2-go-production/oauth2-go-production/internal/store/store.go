package store

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"sync"
	"time"

	"golang.org/x/crypto/argon2"
)

type User struct {
	ID           string
	Username     string
	PasswordHash string
}

type UserStore struct {
	mu    sync.RWMutex
	users map[string]User
}

func NewMemoryUserStore() *UserStore {
	return &UserStore{users: make(map[string]User)}
}

func (s *UserStore) SeedDemoUser(username, password string) error {
	hash, err := HashPassword(password)
	if err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.users[username] = User{ID: "user-123", Username: username, PasswordHash: hash}
	return nil
}

func (s *UserStore) Authenticate(username, password string) (User, bool) {
	s.mu.RLock()
	u, ok := s.users[username]
	s.mu.RUnlock()
	if !ok {
		return User{}, false
	}
	return u, VerifyPassword(u.PasswordHash, password)
}

func HashPassword(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	sum := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	return "argon2id$" + base64.RawStdEncoding.EncodeToString(salt) + "$" +
		base64.RawStdEncoding.EncodeToString(sum), nil
}

func VerifyPassword(encoded, password string) bool {
	var saltB64, hashB64 string
	if _, err := fmtSscanf(encoded, "argon2id$%s$%s", &saltB64, &hashB64); err != nil {
		return false
	}
	salt, err1 := base64.RawStdEncoding.DecodeString(saltB64)
	expected, err2 := base64.RawStdEncoding.DecodeString(hashB64)
	if err1 != nil || err2 != nil {
		return false
	}
	actual := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, uint32(len(expected)))
	return subtleCompare(actual, expected)
}

func fmtSscanf(s, format string, a ...any) (int, error) {
	// tiny parser to avoid adding fmt to the hot path; format is fixed.
	parts := splitDollar(s)
	if len(parts) != 3 || parts[0] != "argon2id" {
		return 0, errors.New("invalid password hash")
	}
	if p, ok := a[0].(*string); ok {
		*p = parts[1]
	}
	if p, ok := a[1].(*string); ok {
		*p = parts[2]
	}
	return 2, nil
}

func splitDollar(s string) []string {
	var out []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '$' {
			out = append(out, s[start:i])
			start = i + 1
		}
	}
	return append(out, s[start:])
}

func subtleCompare(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	var x byte
	for i := range a {
		x |= a[i] ^ b[i]
	}
	return x == 0
}

type AuthorizationCode struct {
	Code          string
	ClientID      string
	UserID        string
	RedirectURI   string
	Scope         string
	CodeChallenge string
	ExpiresAt     time.Time
	Used          bool
}

type AuthorizationCodeStore struct {
	mu    sync.Mutex
	codes map[string]AuthorizationCode
}

func NewMemoryAuthorizationCodeStore() *AuthorizationCodeStore {
	return &AuthorizationCodeStore{codes: make(map[string]AuthorizationCode)}
}

func (s *AuthorizationCodeStore) Put(c AuthorizationCode) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.codes[c.Code] = c
}

func (s *AuthorizationCodeStore) Consume(code string) (AuthorizationCode, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	c, ok := s.codes[code]
	if !ok || c.Used || time.Now().After(c.ExpiresAt) {
		return AuthorizationCode{}, false
	}
	c.Used = true
	s.codes[code] = c
	return c, true
}

type RefreshToken struct {
	Token     string
	UserID    string
	ClientID  string
	Scope     string
	ExpiresAt time.Time
	Used      bool
}

type RefreshTokenStore struct {
	mu     sync.Mutex
	tokens map[string]RefreshToken
}

func NewMemoryRefreshTokenStore() *RefreshTokenStore {
	return &RefreshTokenStore{tokens: make(map[string]RefreshToken)}
}

func (s *RefreshTokenStore) Put(t RefreshToken) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tokens[t.Token] = t
}

func (s *RefreshTokenStore) Rotate(token string) (RefreshToken, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, ok := s.tokens[token]
	if !ok || t.Used || time.Now().After(t.ExpiresAt) {
		return RefreshToken{}, false
	}
	t.Used = true
	s.tokens[token] = t
	return t, true
}

func RandomToken(bytes int) (string, error) {
	b := make([]byte, bytes)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
