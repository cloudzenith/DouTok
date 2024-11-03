package launcher

type App struct {
	Name          string `yaml:"name" json:"name"`
	Version       string `yaml:"version" json:"version"`
	Node          int64  `yaml:"node" json:"node"`
	TraceEndpoint string `json:"trace_endpoint" yaml:"trace_endpoint"`
}
