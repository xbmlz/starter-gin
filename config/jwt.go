package config

type JWT struct {
	Header      string `yaml:"header"`       // header
	SigningKey  string `yaml:"signing-key"`  // jwt签名
	ExpiresTime int    `yaml:"expires-time"` // 过期时间
	RefreshTime int    `yaml:"refresh-time"` // 刷新时间
	Issuer      string `yaml:"issuer"`       // 签发者
}
