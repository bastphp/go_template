package configs

type Elasticsearch struct {
	Host     string `json:"host" yaml:"host" mapstructure:"host"`
	UserName string `mapstructure:"user-name" json:"user-name" yaml:"user-name"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
