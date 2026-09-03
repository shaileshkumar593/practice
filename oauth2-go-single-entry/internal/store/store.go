package store

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"strings"
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

func NewUserStore() *UserStore {
	return &UserStore{users: make(map[string]User)}
}

func (s *UserStore) SeedDemoUser(username, password string) error {
	hash, err := HashPassword(password)
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.users[username] = User{
		ID:           "user-123",
		Username:     username,
		PasswordHash: hash,
	}
	return nil
}

func (s *UserStore) Authenticate(username, password string) (User, bool) {
	s.mu.RLock()
	user, ok := s.users[username]
	s.mu.RUnlock()

	if !ok {
		return User{}, false
	}

	return user, VerifyPassword(user.PasswordHash, password)
}

func HashPassword(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		1,
		64*1024,
		4,
		32,
	)

	return "argon2id$" +
		base64.RawStdEncoding.EncodeToString(salt) +
		"$" +
		base64.RawStdEncoding.EncodeToString(hash), nil
}

func VerifyPassword(encoded, password string) bool {
	parts := strings.Split(encoded, "$")
	if len(parts) != 3 || parts[0] != "argon2id" {
		return false
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false
	}

	expected, err := base64.RawStdEncoding.DecodeString(parts[2])
	if err != nil {
		return false
	}

	actual := argon2.IDKey(
		[]byte(password),
		salt,
		1,
		64*1024,
		4,
		uint32(len(expected)),
	)

	return subtle.ConstantTimeCompare(actual, expected) == 1
}

type Client struct {
	ID           string
	Secret       string
	RedirectURI  string
	AllowedScope string
}

type ClientStore struct {
	mu      sync.RWMutex
	clients map[string]Client
}

func NewClientStore() *ClientStore {
	return &ClientStore{clients: make(map[string]Client)}
}

func (s *ClientStore) Register(client Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[client.ID] = client
}

func (s *ClientStore) Get(id string) (Client, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	client, ok := s.clients[id]
	return client, ok
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

func NewAuthorizationCodeStore() *AuthorizationCodeStore {
	return &AuthorizationCodeStore{codes: make(map[string]AuthorizationCode)}
}

func (s *AuthorizationCodeStore) Put(code AuthorizationCode) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.codes[code.Code] = code
}

func (s *AuthorizationCodeStore) Consume(code string) (AuthorizationCode, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, ok := s.codes[code]
	if !ok || item.Used || time.Now().After(item.ExpiresAt) {
		return AuthorizationCode{}, false
	}

	item.Used = true
	s.codes[code] = item
	return item, true
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

func NewRefreshTokenStore() *RefreshTokenStore {
	return &RefreshTokenStore{tokens: make(map[string]RefreshToken)}
}

func (s *RefreshTokenStore) Put(token RefreshToken) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tokens[token.Token] = token
}

func (s *RefreshTokenStore) Rotate(token string) (RefreshToken, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, ok := s.tokens[token]
	if !ok || item.Used || time.Now().After(item.ExpiresAt) {
		return RefreshToken{}, false
	}

	item.Used = true
	s.tokens[token] = item
	return item, true
}

type RevocationStore struct {
	mu      sync.RWMutex
	revoked map[string]time.Time
}

func NewRevocationStore() *RevocationStore {
	return &RevocationStore{revoked: make(map[string]time.Time)}
}

func (s *RevocationStore) Revoke(token string, expiresAt time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.revoked[token] = expiresAt
}

func (s *RevocationStore) IsRevoked(token string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	expiresAt, ok := s.revoked[token]
	if !ok {
		return false
	}
	return time.Now().Before(expiresAt)
}

func RandomToken(size int) (string, error) {
	if size < 32 {
		return "", errors.New("token size must be >= 32")
	}

	buf := make([]byte, size)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(buf), nil
}
