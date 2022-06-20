package provider

import (
	"encoding/json"
	"io"
	"net/http"
)

func NewKucoinProvider() *KucoinProvider {
	return &KucoinProvider{}
}

type KucoinProvider struct {
}

type MarketStatsResponse struct {
	Code string `json:"code"`
	Data struct {
		Time             int64  `json:"time"`
		Symbol           string `json:"symbol"`
		Buy              string `json:"buy"`
		Sell             string `json:"sell"`
		ChangeRate       string `json:"changeRate"`
		ChangePrice      string `json:"changePrice"`
		High             string `json:"high"`
		Low              string `json:"low"`
		Vol              string `json:"vol"`
		VolValue         string `json:"volValue"`
		Last             string `json:"last"`
		AveragePrice     string `json:"averagePrice"`
		TakerFeeRate     string `json:"takerFeeRate"`
		MakerFeeRate     string `json:"makerFeeRate"`
		TakerCoefficient string `json:"takerCoefficient"`
		MakerCoefficient string `json:"makerCoefficient"`
	} `json:"data"`
}

func (p *KucoinProvider) MarketStatsBTCUSD() (MarketStatsResponse, error) {
	req, err := http.Get("https://api.kucoin.com/api/v1/market/stats?symbol=BTC-USDT")
	all, err := io.ReadAll(req.Body)
	if err != nil {
		return MarketStatsResponse{}, err
	}
	data := MarketStatsResponse{}
	err = json.Unmarshal(all, &data)
	if err != nil {
		return MarketStatsResponse{}, err
	}
	return data, nil
}

