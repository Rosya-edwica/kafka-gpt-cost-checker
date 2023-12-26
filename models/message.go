package models

type Message struct {
	Date       string  `json:"date"`
	UserId     int     `json:"user_id"`
	UserName   string  `json:"user_name"`
	PromtQuery string  `json:"promt"`
	Answer     string  `json:"answer"`
	CostUSD    float32 `json:"cost_usd"`
	CostRUB    float32 `json:"cost_rub"`
	Tokens     int     `json:"tokens"`
	TimeExec   float32 `json:"time_exe"`
}
