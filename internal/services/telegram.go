package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TelegramService struct {
	HTTPClient *http.Client
	ChatId     string
	BotToken   string
}

func NewTelegramService(httpClient *http.Client, botToken string, chatId string) *TelegramService {
	return &TelegramService{
		HTTPClient: httpClient,
		BotToken:   botToken,
		ChatId:     chatId,
	}
}

func (service TelegramService) SendMessage(text string) error {

	requestBody, err := json.Marshal(map[string]string{
		"chat_id": service.ChatId,
		"text":    text,
	})
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", service.BotToken)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := service.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	if res.StatusCode == http.StatusOK {
		fmt.Println("---debug---")
		fmt.Println("ok!")
		fmt.Println("---debug---")
	}
	fmt.Println("---debug---")
	fmt.Println(res.StatusCode)
	fmt.Println("---debug---")

	return nil
}
