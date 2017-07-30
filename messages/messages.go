package messages

import (
	"net/http"

	"io"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// Bot is the Telegram Bot API
type Bot struct {
	Client *tgbotapi.BotAPI
}

// NewBot creates a new bot
func NewBot(telegramKey string) Bot {
	bot, err := tgbotapi.NewBotAPI(telegramKey)
	if err != nil {
		panic(err)
	}
	return Bot{
		Client: bot,
	}
}

// SendMessage sends a message to user
func (b Bot) SendMessage(update tgbotapi.Update, message string) {
	bot := b.Client
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	_, err := bot.Send(msg)
	if err != nil {
		panic(err)
	}
}

// SendNews sends articles on page
func (b Bot) SendNews(update tgbotapi.Update, messages, images []string) {
	bot := b.Client
	for i := 0; i < len(messages); i++ {

		if len(images) > i {
			resp, err := http.Get(images[i])
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			file, err := os.Create("photo.jpg")
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(file, resp.Body)
			if err != nil {
				panic(err)
			}
			file.Close()

			photoConfig := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "photo.jpg")
			tgbotapi.NewDocumentShare(update.Message.Chat.ID, photoConfig.BaseFile.FileID)
			_, err = bot.Send(photoConfig)
			if err != nil {
				panic(err)
			}
		}

		msg := tgbotapi.MessageConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID:           update.Message.Chat.ID,
				ReplyToMessageID: 0,
			},
			Text:                  messages[i],
			ParseMode:             "HTML",
			DisableWebPagePreview: true,
		}

		_, err := bot.Send(msg)
		if err != nil {
			panic(err)
		}
	}
}
