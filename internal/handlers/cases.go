package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func sendCasesMessage(ctx context.Context, b *bot.Bot, update *models.Update, messageText string) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	msg := update.CallbackQuery.Message
	if msg.Message == nil {
		return
	}

	backKeyboard := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "⬅️ Назад", CallbackData: "back"},
			},
		},
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      msg.Message.Chat.ID,
		Text:        messageText,
		ParseMode:   models.ParseModeHTML,
		ReplyMarkup: backKeyboard,
	})
}

func TgCasesHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	sendCasesMessage(ctx, b, update,
		"📄 <b>Кейсы Telegram</b>\n\n<a href='https://telegra.ph/Telegram-Cases-Placeholder'>🔗 Тут будет ссылка на кейс</a>")
}

func VkCasesHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	sendCasesMessage(ctx, b, update,
		"📄 <b>Кейсы ВК</b>\n\n"+
			"<a href='https://telegra.ph/Kak-otdelochniki-20-let-ot-rodu-vlozhili-v-reklamu-800-rublej-i-zarabotali-million-08-11'>🔗 Как отделочники 20 лет от роду вложили в рекламу 800 рублей и заработали миллион</a>\n"+
			"<a href='https://telegra.ph/Kejs-150-000-rub-za-3-dnya-na-internet-magazine-v-VK-08-11'>🔗 150 000 руб. за 3 дня на интернет-магазине в ВК</a>")
}