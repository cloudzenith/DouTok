package etcdx

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Config struct {
	Host    string `json:"host" yaml:"host"`
	Port    int    `json:"port" yaml:"port"`
	Timeout string `json:"timeout" yaml:"timeout"`
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

func (c *Config) GetTimeout() (timeout time.Duration) {
	timeout, err := time.ParseDuration(c.Timeout)
	if err != nil {
		log.Info("etcd timeout parsed error: %v , apply default 3s", c.Timeout)
		timeout = 3 * time.Second
	}
	return
}
