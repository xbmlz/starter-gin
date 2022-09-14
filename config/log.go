package config

// 日志配置
type Log struct {
	Path    string  `yaml:"path"`
	Level   string  `yaml:"level"`
	Encoder string  `yaml:"encoder"`
	Format  string  `yaml:"format"`
	Prefix  string  `yaml:"prefix"`
	Archive Archive `yaml:"archive"`
}

// 归档配置
type Archive struct {
	MaxSize    int  `yaml:"max-size"`
	MaxBackups int  `yaml:"max-backups"`
	MaxAge     int  `yaml:"max-age"`
	Compress   bool `yaml:"compress"`
}
