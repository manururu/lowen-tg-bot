FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bot cmd/bot/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bot .
CMD ["./bot"]