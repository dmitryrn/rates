package service

import "rates/provider"

func NewRatesService(providerCBR *provider.CBRProvider, providerKucoin *provider.KucoinProvider) *RatesService {
	return &RatesService{
		providerCBR:    providerCBR,
		providerKucoin: providerKucoin,
	}
}

type RatesService struct {
	providerCBR    *provider.CBRProvider
	providerKucoin *provider.KucoinProvider
}

func (s *RatesService) Work() error {
	fiatRubRates, err := s.providerCBR.GetRubDailyRates()
	if err != nil {
		return err
	}

	btcUsdRate, err := s.providerKucoin.MarketStatsBTCUSD()
	if err != nil {
		return err
	}


}

func (s *RatesService) Start() error {
	panic("impl")
}

func (s *RatesService) Stop() error {
	panic("impl")
}
