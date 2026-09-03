package token

import (
	"crypto/rsa"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func ParseAndValidate(
	raw string,
	publicKey *rsa.PublicKey,
	issuer string,
) (Claims, error) {
	var claims Claims

	parsed, err := jwt.ParseWithClaims(
		raw,
		&claims,
		func(t *jwt.Token) (any, error) {
			if t.Method != jwt.SigningMethodRS256 {
				return nil, errors.New("unexpected signing method")
			}
			return publicKey, nil
		},
		jwt.WithIssuer(issuer),
		jwt.WithAudience("resource-server"),
		jwt.WithExpirationRequired(),
	)

	if err != nil || !parsed.Valid {
		return Claims{}, errors.New("invalid token")
	}

	return claims, nil
}
