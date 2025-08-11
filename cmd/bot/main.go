package main

import (
	"context"
	"os"

	"github.com/manururu/lowen-tg-bot/internal/telegram"
)

func main() {
	ctx := context.Background()
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		panic("TELEGRAM_TOKEN environment variable is not set")
	}

	botInstance, err := telegram.NewBot(token)
	if err != nil {
		panic(err)
	}

	botInstance.Start(ctx)
}