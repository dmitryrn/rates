package main

import (
	"context"
	"go.uber.org/fx"
	"rates/provider"
	"rates/service"
)

func start(lc fx.Lifecycle, ratesService *service.RatesService) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			println("Starting rates service")
			return ratesService.Start()
		},
		OnStop: func(ctx context.Context) error {
			println("Stopping rates service")
			return ratesService.Stop()
		},
	})
}

func main() {
	app := fx.New(
		fx.Provide(
			service.NewRatesService,
			provider.NewCBRProvider,
			provider.NewKucoinProvider,
		),
		fx.Invoke(start),

		// TODO: use logger
	)

	app.Run()
}
