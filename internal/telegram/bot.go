package telegram

import (
    "context"

	"github.com/manururu/lowen-tg-bot/internal/handlers"

	"github.com/go-telegram/bot"
)

type Bot struct {
	b *bot.Bot
}

func NewBot(token string) (*Bot, error) {
	b, err := bot.New(token)
	if err != nil {
		return nil, err
	}

	tb := &Bot{b: b}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handlers.StartHandler)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "tg_cases", bot.MatchTypeExact, handlers.TgCasesHandler)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "vk_cases", bot.MatchTypeExact, handlers.VkCasesHandler)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "back", bot.MatchTypeExact, handlers.BackHandler)

	return tb, nil
}

func (tb *Bot) Start(ctx context.Context) {
	tb.b.Start(ctx)
}
