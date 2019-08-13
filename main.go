package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

//Config ...
type Config struct {
	Token      string `json :"token"`
	WebhookURL string `json:"webhookURL"`
}

const proxy = "http://94.23.93.151:3128"

func main() {

	file, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}

	currentConfig := Config{}

	err = json.NewDecoder(file).Decode(&currentConfig)

	proxyURL, err := url.Parse(proxy)

	if err != nil {
		fmt.Println(err)
	}

	httpClient := http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}

	bot, err := tgbotapi.NewBotAPIWithClient(currentConfig.Token, &httpClient)

	if err != nil {
		fmt.Println(err)
	}

	bot.Debug = true

	bot.Client.Transport = &http.Transport{
		Proxy: http.ProxyURL(proxyURL)}

	_, err = bot.RemoveWebhook()

	if err != nil {
		fmt.Println(err)
	}
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(currentConfig.WebhookURL))

	info, err := bot.GetWebhookInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)

	if err != nil {
		fmt.Println(err)
	}
	updates := bot.ListenForWebhook("/")

	//for polling method
	/*u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)*/

	go http.ListenAndServe(":8080", nil)
	fmt.Println("server started...")

	for update := range updates {

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text+" "+RandStringRunes(8))
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
