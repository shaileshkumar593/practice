FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:3.22

WORKDIR /app
COPY --from=build /app/server /app/server

EXPOSE 8080

CMD ["/app/server"]
