# OAuth 2.0 in Go — Single Entry Point

A clean, runnable reference implementation of an OAuth 2.0 Authorization Code + PKCE flow in Go.

## Architecture

This version intentionally uses ONE `main.go` and ONE HTTP server:

```text
Browser / OAuth Client
        |
        v
+-----------------------------+
|       Go Application        |
|          :8080              |
|                             |
| /oauth/authorize            |
| /oauth/login                |
| /oauth/token                |
| /oauth/refresh              |
| /oauth/introspect           |
| /oauth/revoke               |
| /.well-known/jwks.json      |
| /.well-known/oauth          |
| /api/profile                |
| /api/orders                 |
| /healthz                    |
+-----------------------------+
```

## Important

This is a production-style educational reference, not a claim that a custom authorization server should replace a mature identity provider without a security review.

The demo uses in-memory stores so it can run with no external database. The project includes PostgreSQL schema scaffolding in `migrations/`.

For production:
- use HTTPS everywhere
- use PostgreSQL/Redis instead of memory stores
- use KMS/HSM/secret manager for signing keys
- implement key rotation and multiple JWKS keys
- add rate limiting, abuse controls and audit logging
- use a mature identity provider when practical
- keep authorization-server authentication/session management hardened
- never log tokens, passwords, codes or private keys

## Prerequisites

- Go 1.24+
- OpenSSL for generating a local RSA key

## Run

```bash
go mod tidy
mkdir -p devkeys
openssl genrsa -out devkeys/private.pem 3072
go run .
```

Server:

```text
http://localhost:8080
```

Demo user:

```text
username: shailesh
password: ChangeMe-123!
```

Demo client:

```text
client_id: demo-client
client_secret: dev-client-secret
redirect_uri: http://localhost:8080/callback
```

## Endpoints

### OAuth endpoints

```text
GET  /oauth/authorize
POST /oauth/authorize
POST /oauth/login
POST /oauth/token
POST /oauth/refresh
POST /oauth/introspect
POST /oauth/revoke
```

### Discovery / keys

```text
GET /.well-known/oauth-authorization-server
GET /.well-known/jwks.json
```

### Protected APIs

```text
GET /api/profile
GET /api/orders
```

### Operational

```text
GET /healthz
GET /
```

## End-to-end browser test

Open:

```text
http://localhost:8080/demo/login
```

The demo client creates:
- `state`
- PKCE `code_verifier`
- PKCE S256 `code_challenge`

Then it redirects to `/oauth/authorize`.

The authorization page accepts:

```text
username: shailesh
password: ChangeMe-123!
```

After login:
1. authorization code is created
2. browser redirects to `/callback`
3. demo client validates `state`
4. client exchanges code + verifier at `/oauth/token`
5. access token is used against `/api/profile`
6. refresh token is available for `/oauth/refresh`

## Manual endpoint tests

### 1. Health

```bash
curl http://localhost:8080/healthz
```

### 2. Discovery

```bash
curl http://localhost:8080/.well-known/oauth-authorization-server
```

### 3. JWKS

```bash
curl http://localhost:8080/.well-known/jwks.json
```

### 4. Protected API without token

```bash
curl -i http://localhost:8080/api/profile
```

Expected:

```text
401 Unauthorized
```

### 5. Full PKCE flow

The easiest way is:

```text
http://localhost:8080/demo/login
```

### 6. Refresh

After the demo login, copy the refresh token and:

```bash
curl -X POST http://localhost:8080/oauth/refresh \
  -H "Content-Type: application/x-www-form-urlencoded" \
  --data-urlencode "refresh_token=YOUR_REFRESH_TOKEN"
```

### 7. Introspection

```bash
curl -X POST http://localhost:8080/oauth/introspect \
  -H "Content-Type: application/x-www-form-urlencoded" \
  --data-urlencode "token=YOUR_ACCESS_TOKEN"
```

### 8. Revoke

```bash
curl -X POST http://localhost:8080/oauth/revoke \
  -H "Content-Type: application/x-www-form-urlencoded" \
  --data-urlencode "token=YOUR_REFRESH_TOKEN"
```

## Token model

Access token:

```text
JWT
  |
  +-- iss
  +-- sub
  +-- aud
  +-- exp
  +-- iat
  +-- jti
  +-- scope
```

Refresh token:

```text
opaque random value
        |
        v
server-side store
```

The refresh token is rotated on use.

## Test with Postman

Create:

```text
GET http://localhost:8080/api/profile
Authorization: Bearer <access-token>
```

The access token can be obtained by using:

```text
http://localhost:8080/demo/login
```

## Production database direction

`migrations/001_oauth.sql` contains tables for:
- clients
- users
- authorization codes
- refresh tokens
- revoked tokens

Replace the in-memory repository implementations under `internal/store/` with PostgreSQL/Redis repositories.
