package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	jsonFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	// Read the file
	jsonFile, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return Config{}, fmt.Errorf("unable to read file: %v", err)
	}

	// Decode the file
	var config Config
	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		return Config{}, fmt.Errorf("unable to Unmarshal JSON: %v", err)
	}

	return config, nil
}

func (cfg Config) SetUser(user string) error {
	cfg, err := Read()
	if err != nil {
		return err
	}
	cfg.CurrentUserName = user
	return write(cfg)
}

func getConfigFilePath() (string, error) {
	// Get the path where the config should be stored
	homeFilePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to find Home directory: %v", err)
	}
	jsonFilePath := homeFilePath + "/" + configFileName
	return jsonFilePath, nil
}

func write(cfg Config) error {
	// Get the path where the config should be stored
	jsonFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Encode the config
	jsonCfg, err := json.Marshal(&cfg)
	if err != nil {
		return fmt.Errorf("unable to marshal JSON: %v", err)
	}

	// Write the file
	err = os.WriteFile(jsonFilePath, jsonCfg, 0644)
	if err != nil {
		return fmt.Errorf("unable to write file: %v", err)
	}
	return nil
}
