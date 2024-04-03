package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rddl-network/rddl-prometheus-exporter/config"

	elements "github.com/rddl-network/elements-rpc"
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

	for _, wallet := range cfg.Wallets {
		setGauge("balance", "Bitcoin balance for network relevant wallet: "+wallet, "elementsd", "wallets", func() float64 {
			res, err := elements.GetBalance(cfg.GetElementsURL(wallet), []string{})
			if err != nil {
				fmt.Printf("%v", err)
				panic(err)
			}
			fmt.Printf("%v", res)
			return res["bitcoin"]
		})
	}

	http.Handle("/metrics", promhttp.Handler())
	logger.Printf("listening on port: %d", cfg.ServicePort)
	log.Fatal(http.ListenAndServe(cfg.GetListenAddr(), nil))
}
