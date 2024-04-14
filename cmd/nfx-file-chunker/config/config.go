package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	S3 struct {
		Region      string `yaml:"region"`
		Url         string `yaml:"url"`
		PartitionId string `yaml:"partitionId"`
	} `yaml:"S3"`
}

func ParseConfig() (*Config, error) {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	return &config, err
}
