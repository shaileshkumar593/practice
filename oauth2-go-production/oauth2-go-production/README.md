# Production-Style OAuth 2.0 Authorization Code + PKCE in Go

This repository is a production-oriented reference implementation for learning and building an OAuth 2.0 Authorization Code + PKCE flow.

## Components

- Authorization Server
  - `/oauth/authorize`
  - `/oauth/token`
  - OIDC-style discovery metadata
  - JWKS endpoint
  - Argon2id password hashing
  - Authorization-code storage abstraction
  - JWT access tokens with RSA signing
  - Refresh-token rotation
  - PKCE S256
  - Exact redirect URI validation
  - Security headers
  - Structured JSON errors
- Resource Server
  - JWT signature/claim validation
  - scope enforcement
  - protected `/api/profile`
- OAuth client
  - Authorization Code + PKCE
  - state validation
  - token exchange
  - protected API call

## Important

This is a production-style reference, not a claim that a custom authorization server is automatically production-ready. For real deployments, prefer a mature identity provider/authorization server unless there is a strong reason to own this security boundary.

## Run

1. Install Go 1.24+.
2. From the project directory:

```bash
go mod tidy
go run ./cmd/auth-server
```

In another terminal:

```bash
go run ./cmd/resource-server
```

In another terminal:

```bash
go run ./cmd/client
```

Open:

```text
http://localhost:8080/login
```

Demo user:

```text
username: shailesh
password: ChangeMe-123!
```

The client redirects to the authorization server at port 9000, then exchanges the code for tokens and calls the resource server at port 9100.

## Environment variables

Auth server:

```text
AUTH_ADDR=:9000
JWT_PRIVATE_KEY_FILE=./devkeys/private.pem
JWT_ISSUER=http://localhost:9000
ACCESS_TOKEN_TTL=10m
REFRESH_TOKEN_TTL=24h
```

Resource server:

```text
RESOURCE_ADDR=:9100
JWT_ISSUER=http://localhost:9000
JWKS_URL=http://localhost:9000/.well-known/jwks.json
```

Client:

```text
CLIENT_ADDR=:8080
AUTH_URL=http://localhost:9000
RESOURCE_URL=http://localhost:9100
CLIENT_ID=demo-web
CLIENT_REDIRECT_URL=http://localhost:8080/callback
```

## Generate development RSA key

```bash
mkdir -p devkeys
openssl genrsa -out devkeys/private.pem 3072
```

Do not commit private keys. The repository `.gitignore` excludes `devkeys/`.

## Production hardening checklist

- Run all endpoints over HTTPS.
- Put secrets and signing keys in a KMS/HSM/secret manager.
- Replace in-memory stores with PostgreSQL/Redis.
- Use a mature identity provider if possible.
- Rotate signing keys and publish a JWKS with overlapping old/new keys.
- Use short-lived access tokens.
- Use refresh-token rotation and reuse detection.
- Add rate limiting and account lockout/abuse controls.
- Add audit logs and security monitoring.
- Add distributed tracing and metrics.
- Validate exact redirect URIs.
- Never log passwords, authorization codes, access tokens, refresh tokens, or private keys.
