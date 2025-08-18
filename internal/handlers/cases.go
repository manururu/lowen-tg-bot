package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func sendCasesMessage(ctx context.Context, b *bot.Bot, update *models.Update, title string, messages []string) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	msg := update.CallbackQuery.Message
	if msg.Message == nil {
		return
	}

	for _, message := range messages {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    msg.Message.Chat.ID,
			Text:      message,
			ParseMode: models.ParseModeHTML,
		})
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
		Text:        title,
		ParseMode:   models.ParseModeHTML,
		ReplyMarkup: backKeyboard,
	})
}

func TgCasesHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	sendCasesMessage(ctx, b, update,
		"📄 <b>Кейсы Telegram</b>",
		[]string{
			"<a href='https://telegra.ph/Telegram-Cases-Placeholder'>🔗 Тут будет ссылка на кейс</a>",
		})
}

func VkCasesHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	sendCasesMessage(ctx, b, update,
		"📄 <b>Кейсы ВК</b>",
		[]string{
			"Ниша: товарный бизнес\nЦель: повысить продажи\nСпособ продвижения: таргет ВК, старый кабинет\n\n<b>Результаты</b>\n\nCTR: 8,3%\nВыручка: 150 тыс. руб.\n\n<a href='https://telegra.ph/Kejs-150-000-rub-za-3-dnya-na-internet-magazine-v-VK-08-11'>🔗 Ссылка на кейс</a>",
			"Ниша: ремонт и строительство\nЦель: привлечь первых клиентов\nСпособ продвижения: таргет ВК на личную страницу\n\n<b>Результаты</b>\n\nЗатраты на рекламу: 2000 руб.\nЛидов привлечено: 3\nВыручка: 310 тыс. руб.\n\n<a href='https://telegra.ph/Kak-otdelochniki-20-let-ot-rodu-vlozhili-v-reklamu-800-rublej-i-zarabotali-million-08-11'>🔗 Ссылка на кейс</a>",
			"Ниша: инфобиз, похудение\nЦель: привлечь лиды на бесплатный интенсив\nСпособ продвижения: VK реклама\n\n<b>Результаты</b>\n\nЗатраты на рекламу: 50 тыс. руб.\nЛидов привлечено: 103\nСтоимость лида: 458 руб.\n\n<a href='https://telegra.ph/Video-VS-foto-prodvizhenie-nutriciologa-v-VK-v-2025-08-13'>🔗 Ссылка на кейс</a>",
		})
}