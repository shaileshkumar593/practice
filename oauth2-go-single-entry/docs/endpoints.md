# Endpoint reference

## GET /oauth/authorize

Parameters:

```text
response_type=code
client_id=demo-client
redirect_uri=http://localhost:8080/callback
scope=profile orders:read
state=<random>
code_challenge=<S256>
code_challenge_method=S256
```

## POST /oauth/login

Demo login form:

```text
client_id
redirect_uri
scope
state
code_challenge
username
password
```

Returns a redirect:

```text
/callback?code=<authorization-code>&state=<state>
```

## POST /oauth/token

Authorization code:

```text
grant_type=authorization_code
code=<code>
client_id=demo-client
client_secret=dev-client-secret
redirect_uri=http://localhost:8080/callback
code_verifier=<verifier>
```

Refresh:

```text
grant_type=refresh_token
client_id=demo-client
client_secret=dev-client-secret
refresh_token=<refresh-token>
```

## POST /oauth/introspect

```text
token=<access-token>
```

## POST /oauth/revoke

```text
token=<token>
```

## GET /api/profile

```http
Authorization: Bearer <access-token>
```

Requires:

```text
profile
```

## GET /api/orders

```http
Authorization: Bearer <access-token>
```

Requires:

```text
orders:read
```
