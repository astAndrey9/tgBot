package main

import (
  "github.com/Syfaro/telegram-bot-api"
  "log"
  "github.com/ramsgoli/Golang-OpenWeatherMap"
  "fmt"
)



func main() {
  	bot, err := tgbotapi.NewBotAPI("1455335552:AAGt66IORSVRDsDfQ9Fm40F5HKFzi-IR_6s")
	owm := openweathermap.OpenWeatherMap{API_KEY: ("068cdf8b46b645699d4334764f7d54cf")}
	  
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
			ChatID := update.Message.Chat.ID
			City := update.Message.Text

			///////
			var currentWeather *openweathermap.CurrentWeatherResponse
			currentWeather, err = owm.CurrentWeatherFromCity(City)
			///////
			log.Printf("[%s] %d %s", UserName, ChatID, City)

			deg := fmt.Sprintf("%.1f", (currentWeather.Main.Temp - 32) / 1.8)
			weather := currentWeather.Name + ": " + deg + " degrees\n" 
			// Созадаем сообщение
			msg := tgbotapi.NewMessage(ChatID, weather)
			// и отправляем его в тг dddd
			bot.Send(msg)
    }
}

