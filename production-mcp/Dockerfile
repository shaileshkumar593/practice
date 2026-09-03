FROM golang:1.25 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" \
    -o /mcp-server ./cmd/mcp-server

FROM gcr.io/distroless/static-debian12

COPY --from=builder /mcp-server /mcp-server

USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/mcp-server"]
