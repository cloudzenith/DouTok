package conf

type Server struct {
	Grpc struct {
		Addr string `yaml:"addr" json:"addr"`
	} `yaml:"grpc" json:"grpc"`
}
