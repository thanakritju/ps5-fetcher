package line

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineService struct {
	bot         *linebot.Client
	recipientId string
}

func NewLineService() *LineService {
	recipientId := os.Getenv("LINE_RECIPIENT_ID")
	bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	return &LineService{bot: bot, recipientId: recipientId}
}

func (s LineService) SendMessage(message string) {
	_, err := s.bot.PushMessage(s.recipientId, linebot.NewTextMessage(message)).Do()
	if err != nil {
		log.Fatal(err)
	}
}
