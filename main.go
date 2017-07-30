package main

import (
	"log"
	"webScraperBot/config"
	"webScraperBot/messages"
	"webScraperBot/scraper"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	config := config.FromYAML("config.yaml")

	//bot := messages.NewBot(os.Getenv("TELEGRAM_KEY"))
	bot := messages.NewBot(config.TelegramKey)

	bot.Client.Debug = true

	log.Printf("Authorized on account %s", bot.Client.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	type lastMessage struct {
		ChatID  int64
		Message string
	}

	//var lastMessages []lastMessage

	updates, _ := bot.Client.GetUpdatesChan(u)

	for update := range updates {

		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.Text == "/start" {

			bot.SendMessage(update, "Hi "+update.Message.Chat.FirstName+", I'm the Al Jazeera web scraper bot, you can send me a message to see if there are new articles to read.\nYou can send the following commands:\n\n/news - get the news\n/help - see all possible commands\n")

		} else if update.Message.Text == "/news" {

			messages, images := scraper.ArticlesScraper()
			bot.SendNews(update, messages, images)

		} else if update.Message.Text == "/help" {

			bot.SendMessage(update, "You can send the following commands:\n\n/news - get the news\n/help - see all possible commands\n")

		} else {

			bot.SendMessage(update, "Sorry, I didn't understand your command. Check out /help if you need to refresh your memory.")

		}

	}
}
