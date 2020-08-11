package lineapi

import (
	"McDailyAutoRun/common"
	"McDailyAutoRun/config"
	"McDailyAutoRun/tool"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

var lineClient *linebot.Client

func init() {
	bot, err := linebot.New(config.LineBotChannelSecret, config.LineBotChannelToken)
	if err != nil {
		fmt.Println("connect line bot fail", err.Error())
	}
	lineClient = bot
}

// EventHandle handle event form line
func EventHandle(w http.ResponseWriter, r *http.Request) {
	events, err := lineClient.ParseRequest(r)
	if err != nil {
		fmt.Println("parse request fail", err.Error())
	}

	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if i := tool.GetUserIndex(event.Source.UserID); i == -1 {
					config.User = append(config.User, common.User{
						LineID: event.Source.UserID,
					})
				}
				userIndex := tool.GetUserIndex(event.Source.UserID)
				prefix := strings.Split(message.Text, ": ")[0]
				if prefix == "McDaily" {
					info := strings.Split(message.Text, ": ")[1]
					config.User[userIndex].McDailyAccount = strings.Split(info, "/")[0]
					config.User[userIndex].McDailyPassword = strings.Split(info, "/")[1]
					if _, err = lineClient.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("新增麥當勞報報成功")).Do(); err != nil {
						fmt.Printf("fail to add McDaily to user: %s", event.Source.UserID)
					}
				}
			}
		case linebot.EventTypeFollow:
			if _, err = lineClient.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("歡迎使用！！")).Do(); err != nil {
				fmt.Printf("fail to reply follow event to user: %s", event.Source.UserID)
			}
		}
	}
}

// GetMessageQuota ...
func GetMessageQuota(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get line bot message quota")

	w.WriteHeader(http.StatusOK)
	result, err := lineClient.GetMessageQuota().Do()
	if err != nil {
		fmt.Println("fail to get message quota")
	}
	json.NewEncoder(w).Encode(result)
}

// PushMessageToUser ...
func PushMessageToUser(ID string, msg string) {
	fmt.Printf("push message: %s to user: %s", msg, ID)

	if _, err := lineClient.PushMessage(ID, linebot.NewTextMessage(msg)).Do(); err != nil {
		fmt.Printf("fail to push message: %s to user: %s", msg, ID)
	}
}
