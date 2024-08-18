package miniox

type Config struct {
	Host        string `yaml:"host" json:"host"`
	Port        int    `yaml:"port" json:"port"`
	ConsolePort int    `yaml:"console_port" json:"console_port"`
	AccessKey   string `yaml:"access_key" json:"access_key"`
	SecretKey   string `yaml:"secret_key" json:"secret_key"`
	Secure      bool   `yaml:"secure" json:"secure"`
}

func (c *Config) SetDefault() {
	if c.Host == "" {
		c.Host = "localhost"
	}

	if c.Port == 0 {
		c.Port = 9000
	}

	if c.ConsolePort == 0 {
		c.ConsolePort = 9001
	}
}
