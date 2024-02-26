package config

import (
	"errors"

	"github.com/spf13/viper"
)

var (
	Server struct {
		Host string
		Port int
	}
)

func Load(cfgPath string) error {

	v := viper.New()
	v.SetConfigFile(cfgPath)

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	// Server
	if err := v.UnmarshalKey("server", &Server); err != nil {
		return errors.New("failed to unmarshal [server] config")
	}

	return nil
}
