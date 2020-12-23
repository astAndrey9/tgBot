package main
import (
  "github.com/Syfaro/telegram-bot-api"
  "log"
)


var apiKey = os.Getenv("068cdf8b46b645699d4334764f7d54cf")

func main() {
  // подключаемся к боту с помощью токена
  bot, err := tgbotapi.NewBotAPI("1455335552:AAGt66IORSVRDsDfQ9Fm40F5HKFzi-IR_6s")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// инициализируем канал, куда будут прилетать обновления от API
	u := tgbotapi.NewUpdate(0)
    u.Timeout = 30
	updates, err := bot.GetUpdatesChan(u)
    // читаем обновления из канала
    for update := range updates {
        			// Пользователь, который написал боту
			UserName := update.Message.From.UserName

			// ID чата/диалога.
			// Может быть идентификатором как чата с пользователем
			// (тогда он равен UserID) так и публичного чата/канала
			ChatID := update.Message.Chat.ID

			// Текст сообщения
			Text := update.Message.Text

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			// Ответим пользователю его же сообщением.
			reply := Text
			// Созадаем сообщение
			msg := tgbotapi.NewMessage(ChatID, reply)
			// и отправляем его
			bot.Send(msg)
    }
}