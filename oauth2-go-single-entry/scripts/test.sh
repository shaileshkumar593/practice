#!/usr/bin/env bash
set -euo pipefail

BASE="http://localhost:8080"

echo "== health =="
curl -s "$BASE/healthz"
echo

echo "== discovery =="
curl -s "$BASE/.well-known/oauth-authorization-server"
echo

echo "== jwks =="
curl -s "$BASE/.well-known/jwks.json"
echo

echo "== protected API without token =="
curl -i -s "$BASE/api/profile" | head
echo
echo "Use http://localhost:8080/demo/login for the complete browser PKCE flow."
