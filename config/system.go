package config

type System struct {
	Env           string `json:"env" yaml:"env"`
	Addr          int    `json:"addr" yaml:"addr"`
	DbType        string `json:"dbType" yaml:"db-type"`
	OssType       string `json:"ossType" yaml:"oss-type"`
	UseMultipoint bool   `json:"useMultipoint" yaml:"use-multipoint"`
}
