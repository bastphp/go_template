package configs

type elasticsearch struct {
	Addr     []string `json:"addr" yaml:"addr" mapstructure:"addr"`
	UserName string   `json:"userName" yaml:"userName" mapstructure:"userName"`
	Password string   `json:"password" yaml:"password" mapstructure:"password"`
}
