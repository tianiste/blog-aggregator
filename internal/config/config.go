package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username

	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %s", err)
	}

	path, err := getConfigLocation()
	if err != nil {
		return fmt.Errorf("error getting config path: %s", err)
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error while writing file: %s", err)
	}

	return nil
}
func Read() Config {
	var c Config

	path, err := getConfigLocation()
	if err != nil {
		fmt.Errorf("error getting config path: %s", err)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return Config{}
	}

	err = json.Unmarshal(file, &c)
	if err != nil {
		fmt.Println(err)
		return Config{}
	}
	return c
}

func getConfigLocation() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("couldn't get user home dir: %s", err)
	}

	configPath := filepath.Join(homedir, configFileName)

	return configPath, nil
}
