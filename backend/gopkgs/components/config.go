package components

type ComponentLoadMap map[string]*ComponentLoadConfig

type ComponentLoadConfig struct {
	Disable   bool   `yaml:"disable" json:"disable" mapstructure:"disable"`
	StoreKey  string `yaml:"storeKey" json:"storeKey" mapstructure:"storeKey"`
	ConfigKey string `yaml:"configKey" json:"configKey" mapstructure:"configKey"`
}
