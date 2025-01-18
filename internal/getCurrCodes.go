package internal

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/jedib0t/go-pretty/v6/table"
	"golang.org/x/net/html/charset"
)

type codes struct {
	XMLName xml.Name `xml:"Valuta"`
	Name    string   `xml:"name,attr"`
	Items   []item   `xml:"Item"`
}

type item struct {
	ID          string `xml:"ID,attr"`
	Name        string `xml:"Name"`
	EngName     string `xml:"EngName"`
	Nominal     int    `xml:"Nominal"`
	ParentCode  string `xml:"ParentCode"`
	ISONumCode  string `xml:"ISO_Num_Code"`
	ISOCharCode string `xml:"ISO_Char_Code"`
}

func GetCurrCodes() {
	var data codes
	c := resty.New()
	response, err := c.R().
		Get("https://cbr.ru/scripts/XML_valFull.asp")
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
	t.AppendHeader(table.Row{"name", "code"})
	for _, v := range data.Items {
		t.AppendRow(table.Row{v.Name, v.ISOCharCode})
	}
	fmt.Println(t.Render())

}
