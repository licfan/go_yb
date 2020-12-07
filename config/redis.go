package config

type Redis struct {
	DB       int    `json:"db111" yaml:"db"`
	Addr     string `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`
}
