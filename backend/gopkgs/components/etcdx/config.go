package etcdx

import "fmt"

type Config struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

func (c *Config) SetDefault() {
	if c.Host == "" {
		c.Host = "localhost"
	}

	if c.Port == 0 {
		c.Port = 2379
	}
}

func (c *Config) GetEntrypoint() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
