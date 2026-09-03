package token

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	"path/filepath"

	"github.com/golang-jwt/jwt/v5"
)

const KeyID = "dev-key-1"

type Claims struct {
	Scope string `json:"scope"`
	jwt.RegisteredClaims
}

func LoadOrGenerateRSAKey(path string) (*rsa.PrivateKey, error) {
	if data, err := os.ReadFile(path); err == nil {
		block, _ := pem.Decode(data)
		if block == nil {
			return nil, errors.New("invalid PEM")
		}

		if key, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
			return key, nil
		}

		value, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}

		key, ok := value.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("private key is not RSA")
		}
		return key, nil
	}

	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return nil, err
	}

	key, err := rsa.GenerateKey(rand.Reader, 3072)
	if err != nil {
		return nil, err
	}

	data := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})

	if err := os.WriteFile(path, data, 0600); err != nil {
		return nil, err
	}

	return key, nil
}

func SignAccessToken(
	key *rsa.PrivateKey,
	issuer string,
	userID string,
	scope string,
	ttlSeconds int64,
) (string, time.Time, error) {
	now := time.Now()
	expires := now.Add(time.Duration(ttlSeconds) * time.Second)

	claims := Claims{
		Scope: scope,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   userID,
			Audience:  jwt.ClaimStrings{"resource-server"},
			ExpiresAt: jwt.NewNumericDate(expires),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        uuid.NewString(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	jwtToken.Header["kid"] = KeyID

	signed, err := jwtToken.SignedString(key)
	return signed, expires, err
}
