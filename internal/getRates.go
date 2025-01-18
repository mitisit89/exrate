package internal

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/jedib0t/go-pretty/v6/table"
	"golang.org/x/net/html/charset"
)

var apiUrl string = "https://cbr.ru/scripts/XML_daily.asp"

// Valute represents each currency in the XML
type valute struct {
	ID        string `xml:"ID,attr"`
	NumCode   int    `xml:"NumCode"`
	CharCode  string `xml:"CharCode"`
	Nominal   int    `xml:"Nominal"`
	Name      string `xml:"Name"`
	Value     string `xml:"Value"`
	VunitRate string `xml:"VunitRate,omitempty"` // Optional field
}

type currency struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []valute `xml:"Valute"`
}

func GetRate(currencyСodes []string) {
	var data currency

	response, err := resty.New().R().
		Get(apiUrl)
	if err != nil {
		fmt.Println(err)
	}
	reader := strings.NewReader(response.String())
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(&data); err != nil {
		fmt.Println(err)
	}
	t := table.NewWriter()
	t.AppendHeader(table.Row{"name", "value"})

	for _, v := range data.Valutes {
		for _, code := range currencyСodes {
			if v.CharCode == strings.ToUpper(code) {
				t.AppendRow(table.Row{v.Name, v.Value})
			}
		}
	}
	fmt.Println("Date: ", data.Date)
	fmt.Println(t.Render())
}
