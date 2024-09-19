package config

import "github.com/spf13/viper"

type HTTP struct {
	Mode string `yaml:"mode"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Server struct {
	HTTP HTTP `yaml:"http"`
}

type Database struct {
	Driver      string `yaml:"driver"`
	DSN         string `yaml:"dsn"`
	AutoMigrate bool   `yaml:"auto_migrate"`
}

type Redis struct {
	Enabled  bool   `yaml:"enabled"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Data struct {
	Database Database `yaml:"database"`
	Redis    Redis    `yaml:"redis"`
}

type Config struct {
	Server Server `yaml:"server"`
	Data   Data   `yaml:"data"`
}

func (h *HTTP) Addr() string {
	return h.Host + ":" + h.Port
}

func Load(configFile string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(configFile)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
