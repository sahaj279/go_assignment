package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type AppConfig struct {
	Database struct {
		User                  string        `yaml:"user"`
		Password              string        `yaml:"password"`
		Host                  string        `yaml:"host"`
		Name                  string        `yaml:"name"`
		MaxIdleConnections    int           `yaml:"maxIdleConnections"`
		MaxOpenConnections    int           `yaml:"maxOpenConnections"`
		MaxConnectionLifeTime time.Duration `yaml:"maxConnectionLifetime"`
		MaxConnectionIdleTime time.Duration `yaml:"maxConnectionIdletime"`
		DisableTLS            bool          `yaml:"disableTLS"`
		Debug                 bool          `yaml:"debug"`
	} `yaml:"database"`
	Consumer struct {
		Count int `yaml:"count"`
	} `yaml:"consumer"`
}

// LoadAppConfig builds config for database and returns a DbConfig struct.
func LoadAppConfig() (conf AppConfig, er error) {
	err := cleanenv.ReadConfig("application.yaml", &conf)
	if err != nil {
		err = errors.Wrap(err, "error occurred while reading application.yaml")
		return AppConfig{}, err
	}
	return conf, nil
}
