package bootstrap

import (
	"encoding/json"
	"fmt"
	"hka-server-login/config"
	"os"
)

// LoadConfig loads the configuration from a JSON file
func LoadConfig(file string) (*config.Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer f.Close()

	var conf config.Config
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&conf); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return &conf, nil
}
