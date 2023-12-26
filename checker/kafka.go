// Читаем здесь сообщения из gpt-cost топика kafka и сохраняем их в JSON и Postgres
package checker

import (
	"context"
	"encoding/json"
	"fmt"
	"kafka_gpt_cost/db"
	"kafka_gpt_cost/models"
	"kafka_gpt_cost/tools"
	"strings"

	"github.com/segmentio/kafka-go"
)

func SaveKafkaMessages() {
	cfg := kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "gpt-cost",
		GroupID:     "my-group",
		StartOffset: kafka.FirstOffset,
	}
	consumer := kafka.NewReader(cfg)
	defer consumer.Close()
	i := 0
	fmt.Println("Слушаем сообщения....")
	for {
		data, err := consumer.ReadMessage(context.Background())
		if err != nil {
			tools.SendErrorToTelegram("Не удалось прочитать сообщение из kafka:" + err.Error())
			break
		}
		message := models.Message{}
		if !strings.Contains(string(data.Value), "answer") {
			fmt.Println("в этом сообщении что-то не то записано:", data.Value)
			continue
		}
		err = json.Unmarshal([]byte(data.Value), &message)
		if err != nil {
			tools.SendErrorToTelegram("Не удалось поместить сообщение в структуру из kafka:" + err.Error())
		}
		i += 1
		fmt.Println("Сообщение №", i)
		db.AddToFile(message)
		db.Add(message)
	}

}
