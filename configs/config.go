package configs

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Xxl    Xxl    `json:"xxl" yaml:"xxl" mapstructure:"xxl"`
}
