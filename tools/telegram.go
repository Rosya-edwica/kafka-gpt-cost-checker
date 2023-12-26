package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const chatId = "544490770"

var TELEGRAM_TOKEN string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	TELEGRAM_TOKEN = os.Getenv("TELEGRAM_TOKEN")
}

func SendErrorToTelegram(message string) {
	url := fmt.Sprintf("%s/sendMessage", fmt.Sprintf("https://api.telegram.org/bot%s", TELEGRAM_TOKEN))
	body, _ := json.Marshal(map[string]string{
		"chat_id": chatId,
		"text":    "Kafka consumer:github.com/Rosya-edwica/kafka-gpt-cost-checker:\n\n" + message,
	})
	response, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
}
