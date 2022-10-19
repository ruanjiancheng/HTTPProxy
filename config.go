package main

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Location struct {
	Pattern     string   `yaml:"pattern"`
	ProxyPass   []string `yaml:"proxy_pass"`
	BalanceMode string   `yaml:"balance_mode"`
}

type Config struct {
	Location            []*Location `yaml:"location"`
	Schema              string      `yaml:"schema"`
	Port                int         `yaml:"port"`
	HealthCheck         bool        `yaml:"tcp_health_check"`
	HealthCheckInterval uint        `yaml:"health_check_interval"`
	MaxAllowed          uint        `yaml:"max_allowed"`
}

func ReadConfig(fileName string) (*Config, error) {
	in, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(in, &config)
	if err != nil {
		return nil, err
	}
	return &config, err
}

func (config *Config) Print() {
	fmt.Printf("Schema: %s\nPort: %d\nHealth Check: %v\nLocation:\n",
		config.Schema, config.Port, config.HealthCheck)
	for _, location := range config.Location {
		fmt.Printf("\tRoute: %s\n\tProxy Pass: %s\n\tMode: %s\n\n",
			location.Pattern, location.ProxyPass, location.BalanceMode)
	}
}

func (config *Config) Vaildation() error {
	if config.Schema != "http" {
		return fmt.Errorf("the schema %s not supported", config.Schema)
	}
	if len(config.Location) == 0 {
		return errors.New("the location cannot be null")
	}
	if config.HealthCheckInterval < 1 {
		return errors.New("health_check_interval must > 0")
	}
	return nil
}
