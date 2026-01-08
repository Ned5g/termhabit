package storage

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	AppName string      `toml:"appname"`
	Styles  StyleConfig `toml:"styles"`
}

type StyleConfig struct {
	Title  string `toml:"title"`
	Cursor string `toml:"cursor"`
	Habit  string `toml:"habit"`
	Done   string `toml:"done"`
	Help   string `toml:"help"`
	Input  string `toml:"input"`
	Border string `toml:"border"`
}

var cfgDefault = Config{
	AppName: "termhabit",
	Styles: StyleConfig{
		Title:  "204",
		Cursor: "204",
		Habit:  "204",
		Done:   "204",
		Help:   "204",
		Input:  "204",
		Border: "63",
	},
}

func InitConfigFile(appName string) (string, error) {
	configDir := filepath.Join(os.Getenv("HOME"), ".config", appName)

	err := os.MkdirAll(configDir, 0o755)
	if err != nil {
		fmt.Printf("failed to create directory: %v", err)
		return "", err
	}

	configPath := filepath.Join(configDir, "config.toml")

	b, err := toml.Marshal(cfgDefault)
	if err != nil {
		fmt.Printf("failed to marshal toml file: %v", err)
		return "", err
	}

	if err := os.WriteFile(configPath, b, 0o644); err != nil {
		fmt.Printf("failed to write config.toml: %v", err)
		return "", err
	}

	return configPath, nil
}

func LoadConfig(appName string) (*Config, error) {
	configPath := filepath.Join(os.Getenv("HOME"), ".config", appName, "config.toml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		InitConfigFile(appName)
	}

	b, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := toml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
