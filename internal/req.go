package internal

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
)

var apiUrl string = "https://www.cbr-xml-daily.ru/daily_json.js"

type currency struct {
	ID       string  `json:"ID"`
	NumCode  string  `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Nominal  int     `json:"Nominal"`
	Name     string  `json:"Name"`
	Value    float64 `json:"Value"`
	Previous float64 `json:"Previous"`
}

type data struct {
	Date         string              `json:"Date"`
	PreviousDate string              `json:"PreviousDate"`
	PreviousURL  string              `json:"PreviousURL"`
	Timestamp    string              `json:"Timestamp"`
	Valute       map[string]currency `json:"Valute"`
}

func GetRate(currencyСode string) {
	var data data
	err := requests.
		URL(apiUrl).
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data.Valute[currencyСode])
}
