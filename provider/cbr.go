package provider

import (
	"encoding/xml"
	"io"
	"net/http"
)

func NewCBRProvider() *CBRProvider {
	return &CBRProvider{}
}

type CBRProvider struct {
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Text    string   `xml:",chardata"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valute  []struct {
		Text     string `xml:",chardata"`
		ID       string `xml:"ID,attr"`
		NumCode  string `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Nominal  string `xml:"Nominal"`
		Name     string `xml:"Name"`
		Value    string `xml:"Value"`
	} `xml:"Valute"`
}

func (p CBRProvider) GetRubDailyRates() (ValCurs, error) {
	req, err := http.Get("https://www.cbr.ru/scripts/XML_daily.asp")
	all, err := io.ReadAll(req.Body)
	if err != nil {
		return ValCurs{}, err
	}
	data := ValCurs{}
	err = xml.Unmarshal(all, &data)
	if err != nil {
		return ValCurs{}, err
	}
	return data, nil
}

