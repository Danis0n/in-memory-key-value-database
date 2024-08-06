package configuration

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

var (
	errNoFilename = errors.New("filename invalid")
)

type Configuration struct {
	EngineConfiguration  *EngineConfiguration  `yaml:"engine"`
	NetworkConfiguration *NetworkConfiguration `yaml:"network"`
	LoggingConfiguration *LoggingConfiguration `yaml:"logging"`
	RunnerConfiguration  *RunnerConfiguration  `yaml:"runner"`
}

type RunnerConfiguration struct {
	Type string `yaml:"type"`
}

type LoggingConfiguration struct {
	Level  string `yaml:"level"`
	Output string `yaml:"output"`
}

type EngineConfiguration struct {
	Type string `yaml:"type"`
}

type NetworkConfiguration struct {
	Address        string        `yaml:"address"`
	MaxConnections int           `yaml:"max_connections"`
	MaxMessageSize string        `yaml:"max_message_size"`
	IdleTimeout    time.Duration `yaml:"idle_timeout"`
}

func Load(configPath string) (*Configuration, error) {

	if configPath == "" {
		return &Configuration{}, errNoFilename
	}

	var config *Configuration
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
