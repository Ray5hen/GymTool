package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/line/line-bot-sdk-go/linebot"
	//"line/line-bot-sdk-go/linebot"
	gymtool "github.com/ray5hen/gymtool/tools"
)
var bot *linebot.Client
func main() {
	log.Println("Hello From Go")
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				   //kg, err := strconv.Atoi(message.Text) * 
   				   //if err != nil {
      				// handle error
   				   //}
				var a = gymtool.Gt("p111")
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text+a)).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.StickerMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("I can't read Sticker")).Do(); err != nil {
					log.Print(err)
				}

			}
		}
	}
}