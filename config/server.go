package config

// 服务配置
type Server struct {
	Mode    string `yaml:"mode"`
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}
