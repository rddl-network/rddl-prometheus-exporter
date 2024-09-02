package main

import (
	"context"
	"log"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	elementsrpc "github.com/rddl-network/elements-rpc"
	"github.com/rddl-network/rddl-prometheus-exporter/config"
	"github.com/rddl-network/rddl-prometheus-exporter/elements"
	"github.com/rddl-network/rddl-prometheus-exporter/planetmint"
	"github.com/rddl-network/rddl-prometheus-exporter/system"
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

func registerGauges(ctx context.Context, logger *log.Logger, cfg *config.Config) {
	units := cfg.GetServiceUnits()
	for _, unit := range units {
		unit := unit
		logger.Print("registering gauge for service unit: " + unit)
		unitName := strings.Split(unit, ".")[0]
		sanitizedUnit := strings.ReplaceAll(unitName, "-", "_")
		setGauge("service_active_state_"+sanitizedUnit, "ActiveState for network relevant service: "+unitName, "systemd", "units", func() float64 {
			return system.CheckUnitActiveState(ctx, unit)
		})
	}

	wallets, err := elementsrpc.ListWallets(cfg.GetElementsURL(""), []string{})
	if err != nil {
		log.Fatalf("fatal error fetching wallets: %s", err.Error())
	}

	for _, wallet := range wallets {
		wallet := strings.TrimSpace(wallet)
		sanitizedWallet := strings.ReplaceAll(wallet, "-", "_")
		logger.Print("registering gauge for wallet: " + wallet)
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

	logger.Print("registering gauge for active devices")
	setGauge("active_count", "Devices actively connected to MQTT for PoP participation: ", "planetmint_god", "devices", func() float64 {
		count, err := planetmint.GetActiveDeviceCount()
		if err != nil {
			log.Print(err.Error())
			return 0
		}
		return count
	})

	logger.Print("registering gauge for activated devices")
	setGauge("activated_count", "Devices activated on Planetmint for use on network: ", "planetmint_god", "devices", func() float64 {
		count, err := planetmint.GetActivatedDeviceCount()
		if err != nil {
			log.Print(err.Error())
			return 0
		}
		return count
	})
}
