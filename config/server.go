package config

// 服务配置
type Server struct {
	RunMode string `yaml:"run-mode"`
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}
