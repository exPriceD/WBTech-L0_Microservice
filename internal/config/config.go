package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Host         string `yaml:"host"`
		Port         string `yaml:"port"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		DatabaseName string `yaml:"name"`
	} `yaml:"database"`

	NATS struct {
		Url       string `yaml:"host"`
		ClusterID string `yaml:"cluster_id"`
		ClientID  string `yaml:"client_id"`
	}
}

func LoadConfig() (*Config, error) {
	configPath := filepath.Join("configs", "config.yaml")

	yamlData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var config Config

	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
