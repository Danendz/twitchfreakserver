package entities

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Paths map[string]string `json:"paths"`
}

var GlobalConfig *Config

func NewConfig() error {
	config := &Config{
		Paths: map[string]string{
			"data": "./data",
		},
	}

	file, _ := json.Marshal(config)

	if err := os.WriteFile("config.json", file, 0644); err != nil {
		return err
	}

	GlobalConfig = config

	return nil
}

func SetConfigFromFile(configFile *os.File) error {
	configBytes, _ := io.ReadAll(configFile)

	return json.Unmarshal([]byte(configBytes), &GlobalConfig)
}
