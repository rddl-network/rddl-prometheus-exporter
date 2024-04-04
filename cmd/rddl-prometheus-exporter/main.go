package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rddl-network/rddl-prometheus-exporter/config"
	"github.com/rddl-network/rddl-prometheus-exporter/elements"
)

func setGauge(name string, help string, namespace string, subsystem string, callback func() float64) {
	gaugeFunc := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      name,
		Help:      help,
	}, callback)
	prometheus.MustRegister(gaugeFunc)
}

func main() {
	cfg, err := config.LoadConfig("./")
	if err != nil {
		log.Fatalf("fatal error loading config file: %s", err.Error())
	}

	logger := log.Default()

	for _, wallet := range strings.Split(cfg.Wallets, ",") {
		wallet := strings.TrimSpace(wallet)
		sanitizedWallet := strings.ReplaceAll(wallet, "-", "_")
		logger.Printf("registering gauge for wallet: " + wallet)
		setGauge("balance_"+sanitizedWallet, "Bitcoin balance for network relevant wallet: "+wallet, "elementsd", "wallets", func() float64 {
			url := cfg.GetElementsURL(wallet)
			balance, err := elements.GetWalletBalance(url, wallet)
			if err != nil {
				log.Print(err.Error())
				return 0
			}
			return balance
		})
	}

	http.Handle("/metrics", promhttp.Handler())
	logger.Printf("listening on port: %d", cfg.ServicePort)
	log.Fatal(http.ListenAndServe(cfg.GetListenAddr(), nil))
}
