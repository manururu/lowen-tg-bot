FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bot cmd/bot/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bot .
CMD ["./bot"]
