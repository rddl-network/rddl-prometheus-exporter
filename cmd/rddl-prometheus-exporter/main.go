package main

import (
	"context"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rddl-network/rddl-prometheus-exporter/config"
)

func main() {
	cfg, err := config.LoadConfig("./")
	if err != nil {
		log.Fatalf("fatal error loading config file: %s", err.Error())
	}

	logger := log.Default()

	registerGauges(context.Background(), logger, cfg)

	http.Handle("/metrics", promhttp.Handler())
	logger.Printf("listening on port: %d", cfg.ServicePort)
	log.Fatal(http.ListenAndServe(cfg.GetListenAddr(), nil))
}
