package config

// 全局配置
type Config struct {
	Server     Server
	Log        Log
	Redis      Redis
	Datasource Datasource
	JWT        JWT
}
