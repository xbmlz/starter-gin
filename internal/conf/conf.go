package conf

import (
	"errors"

	"github.com/spf13/viper"
)

type Server struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Log struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	Compress   bool   `mapstructure:"compress"`
}

type Database struct {
	Type string `mapstructure:"type"`
	DSN  string `mapstructure:"dsn"`
}

type JWT struct {
	Secret string `mapstructure:"secret"`
	Expire int    `mapstructure:"expire"`
}

type Security struct {
	JWTSecretKey string `mapstructure:"jwt_secret_key"`
	JWTExpire    int    `mapstructure:"jwt_expire"`
}

type AllConfig struct {
	Server   Server   `mapstructure:"server"`
	Log      Log      `mapstructure:"log"`
	Database Database `mapstructure:"database"`
	Security Security `mapstructure:"security"`
}

var Config AllConfig

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

	if err := v.Unmarshal(&Config); err != nil {
		return errors.New("failed to unmarshal config")
	}
	return nil
}
