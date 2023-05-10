package configs

type System struct {
	Addr          int `json:"addr" yaml:"addr" mapstructure:"addr"`
	ReaderTimeOut int `json:"reader_time_out" yaml:"reader_time_out" mapstructure:"reader_time_out"`
	WriterTimeOut int `json:"writer_time_out" yaml:"writer_time_out" mapstructure:"writer_time_out"`
	MaxParams     int `json:"max_params" yaml:"max_params" mapstructure:"max_params"`
}
