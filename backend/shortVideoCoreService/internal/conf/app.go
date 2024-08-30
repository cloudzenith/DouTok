package conf

type App struct {
	Name    string `yaml:"name" json:"name"`
	Version string `yaml:"version" json:"version"`
	Node    int64  `yaml:"node" json:"node"`
}
