package config

// 数据源
type Datasource struct {
	Primary string `yaml:"primary"`
	Sqlite  Sqlite
}

// sqlite
type Sqlite struct {
	Path string `yaml:"path"`
}
