package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

// API provides basic telegram functionality
type API struct {
	chatID  int64
	sendURI string
}

// NewAPI instantiates api
func NewAPI(chatID int64, botToken string) *API {
	sendURI := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	log.Info().Msgf("Telegram uri is:" + sendURI)
	return &API{chatID, sendURI}
}

type reqBody struct {
	ChatID    int64  `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

// SendMessage ...
func (api *API) SendMessage(text string) {
	// odd numbers of "_" break Telegram parser in "markdown" mode
	// so let's workarund it
	text = strings.Replace(text, "_", "\\_", -1)

	req := &reqBody{
		ChatID:    api.chatID,
		Text:      text,
		ParseMode: "markdown",
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		log.Error().Err(err).Msg("Error marshling json")
	}

	log.Info().Msgf("Sending Message: '%s'\n", text)

	res, err := http.Post(api.sendURI, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		log.Error().Err(err).Msg("Error posting")
	}

	if res.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(res.Body)
		bodyStr := ""
		if err == nil {
			bodyStr = string(body)
		}
		log.Error().Msgf("Send Message: unexpected status: %v\nResp:%v", res.Status, bodyStr)
	}

	log.Info().Msg("Message sent")
}
