package launcher

type App struct {
	Name          string `json:"name" yaml:"name"`
	Version       string `json:"version" yaml:"version"`
	Node          int64  `json:"node" yaml:"node"`
	TempoEndpoint string `json:"tempo_endpoint" yaml:"tempo_endpoint"`
}
