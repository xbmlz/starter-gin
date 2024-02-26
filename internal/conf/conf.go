package conf

import (
	"errors"

	"github.com/spf13/viper"
)

var (
	Server struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	}

	Log struct {
		Level      string `mapstructure:"level"`
		Filename   string `mapstructure:"filename"`
		MaxSize    int    `mapstructure:"max_size"`
		MaxAge     int    `mapstructure:"max_age"`
		MaxBackups int    `mapstructure:"max_backups"`
		Compress   bool   `mapstructure:"compress"`
	}
)

func Load(customConf string) error {
	v := viper.New()

	if customConf != "" {
		customConf = "config.yaml"
	}

	v.SetConfigFile(customConf)
	if err := v.ReadInConfig(); err != nil {
		return errors.New("failed to read config file")
	}
	v.SetConfigType("yaml")

	// Server
	if err := v.UnmarshalKey("server", &Server); err != nil {
		return errors.New("failed to unmarshal [server] config")
	}

	// Log
	if err := v.UnmarshalKey("log", &Log); err != nil {
		return errors.New("failed to unmarshal [log] config")
	}

	return nil
}
