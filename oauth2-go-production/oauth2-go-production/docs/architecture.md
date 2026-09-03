# Architecture

```text
Browser
   |
   | Authorization Code + PKCE
   v
Authorization Server
   |   | \-- PostgreSQL / Redis (production)
   |
   +----> JWKS
   |
   v
Resource Server
   |
   +----> protected APIs
```

## Security boundaries

1. Browser authentication occurs at the authorization server.
2. Authorization code is short-lived and one-time use.
3. PKCE S256 binds the code exchange to the client transaction.
4. Access tokens are short-lived JWTs.
5. Refresh tokens are opaque, rotated, and stored server-side.
6. Resource server validates issuer, audience, signature, expiry and scopes.
7. Redirect URIs are exact-match validated.

## Production replacement points

The demo uses memory stores and a local RSA key only to make the flow runnable. Replace them with PostgreSQL/Redis and KMS/HSM-backed signing key management before production.
