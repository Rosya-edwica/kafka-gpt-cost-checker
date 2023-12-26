package db

import (
	"encoding/json"
	"io/ioutil"
	"kafka_gpt_cost/models"
	"kafka_gpt_cost/tools"
	"os"
)

type Data struct {
	Items []models.Message `json:"values"`
}

func AddToFile(message models.Message) {
	file, err := os.OpenFile("db.json", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		tools.SendErrorToTelegram("Не удалось добавить значение в файл JSON:" + err.Error())
	}
	fileData, err := ioutil.ReadAll(file)

	var data Data
	json.Unmarshal(fileData, &data)
	file.Close()

	data.Items = append(data.Items, message)

	fileW, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		tools.SendErrorToTelegram("Не удалось добавить значение в файл JSON:" + err.Error())
	}
	err = ioutil.WriteFile("db.json", fileW, 0644)
	if err != nil {
		tools.SendErrorToTelegram("Не удалось добавить значение в файл JSON:" + err.Error())
	}
}
