package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func MainMenu(ctx context.Context, b *bot.Bot, chatID int64) {
	keyboard := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Кейсы TG", CallbackData: "tg_cases"},
			},
			{
				{Text: "Кейсы ВК", CallbackData: "vk_cases"},
			},
		},
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        "Выберите:",
		ReplyMarkup: keyboard,
	})
}