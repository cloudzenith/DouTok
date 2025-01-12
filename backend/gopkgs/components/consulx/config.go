package consulx

type Config struct {
	Address string `yaml:"address" json:"address" `
}

func (c *Config) SetDefault() {
	if c.Address == "" {
		c.Address = "127.0.0.1:8500"
	}
}
