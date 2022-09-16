package config

type Redis struct {
	Enable   bool   `yaml:"enable"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database int    `yaml:"database"`
	Password string `yaml:"password"`
}
