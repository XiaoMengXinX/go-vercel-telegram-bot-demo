package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Msg    string `json:"text"`
	ChatID int64  `json:"chat_id"`
	Method string `json:"method"`
}

func Repeater2(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)

	var update tgbotapi.Update

	err := json.Unmarshal(body, &update)
	if err != nil {
		log.Println(err)
		return
	}

	if update.Message.Text != "" {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		data := Response{
			Msg:    update.Message.Text,
			Method: "sendMessage",
			ChatID: update.Message.Chat.ID,
		}
		msg, _ := json.Marshal(data)

		log.Printf("Response %s", string(msg))

		w.Header().Add("Content-Type", "application/json")

		fmt.Fprintf(w, string(msg))
	}
}
