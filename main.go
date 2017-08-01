package main

import (
	"log"
	"strconv"
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

	type LastMessage struct {
		ChatID   int64
		Message  string
		Articles []string
	}

	var lastMessages []messages.LastMessage

	updates, _ := bot.Client.GetUpdatesChan(u)

	for update := range updates {

		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.Text == "/start" {

			bot.SendMessage(update, "Hi "+update.Message.Chat.FirstName+", I'm the Al Jazeera web scraper bot, you can send me a message to see if there are new articles to read.\nYou can send the following commands:\n\n/news - get the news\n/help - see all possible commands\n")

		} else if update.Message.Text == "/news" {

			articles, images := scraper.ArticlesScraper()
			lastMessages = append(
				lastMessages,
				messages.LastMessage{
					ChatID:   update.Message.Chat.ID,
					Articles: articles,
				})
			bot.SendNews(update, articles, images)

		} else if update.Message.Text == "Yes" {

			for _, lastMessage := range lastMessages {
				if update.Message.Chat.ID == lastMessage.ChatID {
					bot.SendKeyboard(update, len(lastMessage.Articles))
				}
			}

		} else if update.Message.Text == "1" || update.Message.Text == "2" || update.Message.Text == "3" || update.Message.Text == "4" || update.Message.Text == "5" {

			articleNumber, err := strconv.Atoi(update.Message.Text)
			if err != nil {
				panic(err)
			}
			for _, lastMessage := range lastMessages {
				if update.Message.Chat.ID == lastMessage.ChatID {
					for i := 0; i < 5; i++ {
						log.Printf("Article %v is: %v", i, lastMessage.Articles[i])
					}
					log.Printf("Article number is: %v\n", articleNumber)
					article := scraper.ArticleScraper(lastMessage.Articles[articleNumber-1])
					bot.SendMessage(update, article)
				}
			}
			// could create lastMessage var at the beginning, before checking the update.Message.Text
			// but problem if it is first time, not initialized

		} else if update.Message.Text == "/help" {

			bot.SendMessage(update, "You can send the following commands:\n\n/news - get the news\n/help - see all possible commands\n")

		} else {

			bot.SendMessage(update, "Sorry, I didn't understand your command. Check out /help if you need to refresh your memory.")

		}

	}
}
