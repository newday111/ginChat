package conf

import (
	"os"

	"github.com/goccy/go-yaml"
)

type ConfigStruct struct {
	App struct {
		Env     string `yaml:"env"`
		LogPath string `yaml:"logPath"`
	} `yaml:"app"`
}

func LoadConfig(configFilePath string) (*ConfigStruct, error) {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var config ConfigStruct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
