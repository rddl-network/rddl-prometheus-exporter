package config

import (
	"fmt"
	"strings"
	"sync"
)

const DefaultConfigTemplate = `
rpc-host="{{ .RPCHost }}"
rpc-user="{{ .RPCUser }}"
rpc-pass="{{ .RPCPass }}"
service-bind="{{ .ServiceBind }}"
service-port={{ .ServicePort }}
service-units="{{ .ServiceUnits }}"
planetmint-host="{{ .PlanetmintHost }}"
`

type Config struct {
	RPCHost        string `mapstructure:"rpc-host"`
	RPCUser        string `mapstructure:"rpc-user"`
	RPCPass        string `mapstructure:"rpc-pass"`
	ServiceBind    string `mapstructure:"service-bind"`
	ServicePort    int    `mapstructure:"service-port"`
	ServiceUnits   string `mapstructure:"service-units"`
	PlanetmintHost string `mapstructure:"planetmint-host"`
}

var (
	config     *Config
	initConfig sync.Once
)

// DefaultConfig returns distribution-service default config
func DefaultConfig() *Config {
	return &Config{
		RPCHost:        "localhost:18884",
		RPCUser:        "user",
		RPCPass:        "password",
		ServiceBind:    "localhost",
		ServicePort:    8080,
		ServiceUnits:   "",
		PlanetmintHost: "127.0.0.1:9090",
	}
}

func GetConfig() *Config {
	initConfig.Do(func() {
		config = DefaultConfig()
	})
	return config
}

func (c *Config) GetListenAddr() string {
	return fmt.Sprintf("%s:%d", c.ServiceBind, c.ServicePort)
}

func (c *Config) GetElementsURL(wallet string) string {
	url := fmt.Sprintf("http://%s:%s@%s/wallet/%s", c.RPCUser, c.RPCPass, c.RPCHost, wallet)
	return url
}

func (c *Config) GetServiceUnits() []string {
	return strings.Split(c.ServiceUnits, ",")
}
