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
			"<i>Ниша:</i> аналитика для WB\n<i>Цель:</i> протестировать заходы, найти лучшую связку\n<i>Способ продвижения:</i> ТГ АДС\n\n<b>Результаты</b>\n\nНайдены заходы с CTR 0,63 и 0,93.\nКонверсия в подписку на канал - 44%\n\nПродолжаем сотрудничество с заказчиком, масштабируем связку и ведем работу по поиску новых.\n\n<a href='https://telegra.ph/Test-dlya-sistemy-analitiki-WB-vyvody-i-rezultaty-08-20'>🔗 Ссылка на кейс</a>",
		})
}

func VkCasesHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	sendCasesMessage(ctx, b, update,
		"📄 <b>Кейсы ВК</b>",
		[]string{
			"<i>Ниша:</i> товарный бизнес\n<i>Цель:</i> повысить продажи\n<i>Способ продвижения:</i> таргет ВК, старый кабинет\n\n<b>Результаты</b>\n\nCTR: 8,3%\nВыручка: 150 тыс. руб.\n\n<a href='https://telegra.ph/Kejs-150-000-rub-za-3-dnya-na-internet-magazine-v-VK-08-11'>🔗 Ссылка на кейс</a>",
			"<i>Ниша:</i> ремонт и строительство\n<i>Цель:</i> привлечь первых клиентов\n<i>Способ продвижения:</i> таргет ВК на личную страницу\n\n<b>Результаты</b>\n\nЗатраты на рекламу: 2000 руб.\nЛидов привлечено: 3\nВыручка: 310 тыс. руб.\n\n<a href='https://telegra.ph/Kak-otdelochniki-20-let-ot-rodu-vlozhili-v-reklamu-800-rublej-i-zarabotali-million-08-11'>🔗 Ссылка на кейс</a>",
			"<i>Ниша:</i> инфобиз, похудение\n<i>Цель:</i> привлечь лиды на бесплатный интенсив\n<i>Способ продвижения:</i> VK реклама\n\n<b>Результаты</b>\n\nЗатраты на рекламу: 50 тыс. руб.\nЛидов привлечено: 103\nСтоимость лида: 458 руб.\n\n<a href='https://telegra.ph/Video-VS-foto-prodvizhenie-nutriciologa-v-VK-v-2025-08-13'>🔗 Ссылка на кейс</a>",
		})
}