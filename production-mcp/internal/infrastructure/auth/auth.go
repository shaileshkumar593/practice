package auth

import (
    "context"
    "fmt"
)

type Principal struct {
    UserID string
    Roles  []string
}

type Validator interface {
    Validate(ctx context.Context, token string) (*Principal, error)
}

type JWTValidator struct{}

func NewJWTValidator() *JWTValidator {
    return &JWTValidator{}
}

func (v *JWTValidator) Validate(ctx context.Context, token string) (*Principal, error) {
    if token == "" {
        return nil, fmt.Errorf("missing authorization token")
    }

    // Replace with real OIDC/JWT validation in production.
    return &Principal{
        UserID: "example-user",
        Roles:  []string{"knowledge-reader"},
    }, nil
}
