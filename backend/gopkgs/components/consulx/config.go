package consulx

type Config struct {
	Address string `yaml:"address" json:"address" `
}

func (c *Config) SetDefault() {

}
