package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

//Config ...
type Config struct {
	Token string `json :"token"`
}

func main() {
	file, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}

	currentConfig := Config{}

	err = json.NewDecoder(file).Decode(&currentConfig)

	bot, err := tgbotapi.NewBotAPI(currentConfig.Token)

	if err != nil {
		fmt.Println(err)
	}

	go http.ListenAndServe("127.0.0.1:8080", nil)
	updates := bot.ListenForWebhook("/" + bot.Token)

	for update := range updates {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text+RandStringRunes(8))
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)

	}

}

//RandStringRunes for generating random word
func RandStringRunes(n int) string {

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//set webhook string

//https://api.telegram.org/bot972295397:AAEiO9wfDNVd1ec_M6dI1EW8ZMCfGyabW_w/setWebhook?url=https://a9ed699d.ngrok.io/972295397:AAEiO9wfDNVd1ec_M6dI1EW8ZMCfGyabW_w
