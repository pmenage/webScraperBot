package messages

import (
	"net/http"
	"strconv"

	"io"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// Bot is the Telegram Bot API
type Bot struct {
	Client *tgbotapi.BotAPI
}

// LastMessage contains the last articles
type LastMessage struct {
	ChatID   int64
	Message  string
	Articles []string
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

// SetMessage sets a new message
func (l *LastMessage) SetMessage(message string) {
	l.Message = message
}

// SetArticles sets the last articles
func (l *LastMessage) SetArticles(articles []string) {
	l.Articles = articles
}

// SendMessage sends a message to user
func (b Bot) SendMessage(update tgbotapi.Update, message string) {
	bot := b.Client
	msg := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           update.Message.Chat.ID,
			ReplyToMessageID: 0,
		},
		Text:                  message,
		ParseMode:             "HTML",
		DisableWebPagePreview: true,
	}

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

		b.SendMessage(update, messages[i])
	}

	b.SendMessage(update, "Do you want to read an article?")

}

// SendKeyboard sends a keyboard to chose an article
func (b Bot) SendKeyboard(update tgbotapi.Update, numberArticles int) {
	bot := b.Client
	msg := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"Please choose which article you wish to read.")

	var keyboard [][]tgbotapi.KeyboardButton

	for i := 1; i < numberArticles+1; i++ {
		keyboard = append(
			keyboard,
			[]tgbotapi.KeyboardButton{
				tgbotapi.NewKeyboardButton(strconv.Itoa(i)),
			})
	}

	markup := tgbotapi.ReplyKeyboardMarkup{
		Keyboard:        keyboard,
		OneTimeKeyboard: true,
	}
	msg.ReplyMarkup = markup
	_, err := bot.Send(msg)
	if err != nil {
		panic(err)
	}
}
