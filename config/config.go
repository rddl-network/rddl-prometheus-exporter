package config

import (
	"fmt"
	"sync"
)

const DefaultConfigTemplate = `
wallet="{{ .Wallets }}"
rpc-host="{{ .RPCHost }}"
rpc-user="{{ .RPCUser }}"
rpc-pass="{{ .RPCPass }}"
service-bind="{{ .ServiceBind }}"
service-port="{{ .ServicePort }}"
`

type Config struct {
	Wallets     []string `mapstructure:"wallets"`
	RPCHost     string   `mapstructure:"rpc-host"`
	RPCUser     string   `mapstructure:"rpc-user"`
	RPCPass     string   `mapstructure:"rpc-pass"`
	ServiceBind string   `mapstructure:"service-bind"`
	ServicePort int      `mapstructure:"service-port"`
}

var (
	config     *Config
	initConfig sync.Once
)

// DefaultConfig returns distribution-service default config
func DefaultConfig() *Config {
	return &Config{
		Wallets:     []string{"dao", "pop", "issuerwallet", "early-investor-wallet"},
		RPCHost:     "localhost:18884",
		RPCUser:     "user",
		RPCPass:     "password",
		ServiceBind: "localhost",
		ServicePort: 8080,
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
