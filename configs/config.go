package configs

type Server struct {
	System        System        `mapstructure:"system" json:"system" yaml:"system"`
	Elasticsearch elasticsearch `json:"elasticsearch" mapstructure:"system" yaml:"system"`
}
