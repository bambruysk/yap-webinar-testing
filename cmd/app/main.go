package main

import (
	"time"

	"go.uber.org/zap"

	"webinar-testing/internal/api/http"
	"webinar-testing/internal/connectors/warehouse"
	"webinar-testing/internal/service/cart"
	"webinar-testing/internal/storage/inmemory"
)

func main() {
	logger := zap.Must(zap.NewDevelopment())

	wh := warehouse.New(&warehouse.Options{Addr: ":9090"})

	store := inmemory.New()
	usecase := cart.New(logger, store, wh)

	server := http.New(
		&http.Options{Addr: ":8080"}, logger, usecase)

	server.Run()

	logger.Info("start http server")
	time.Sleep(1000 * time.Second)
}
