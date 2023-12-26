// здесь мы сохраняем данные в postgres, чтобы другие программы могли подключиться к нему и прочитать оттуда данные
// на 26 декабря 2023 используется тестовый режим, где postgres запускается локально на сервере, где распложены все программы
// использующие kafka и postgres

package db

import (
	"database/sql"
	"fmt"
	"kafka_gpt_cost/models"
	"kafka_gpt_cost/tools"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "edwica"
	password = "hello_test"
	dbname   = "edwica_germany"
)

func Add(message models.Message) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		tools.SendErrorToTelegram("Не удалось добавить значение в Postgres:" + err.Error())
	}
	stmt, err := db.Prepare(`INSERT INTO gpt_history(
		date, tg_user_id, tg_user_name,
		promt, answer, cost_usd,
		cost_rub, tokens, time_ex
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`)

	_, err = stmt.Exec(message.Date, message.UserId, message.UserName,
		message.PromtQuery, message.Answer, message.CostUSD, message.CostRUB,
		message.Tokens, message.TimeExec)
	db.Close()

}
