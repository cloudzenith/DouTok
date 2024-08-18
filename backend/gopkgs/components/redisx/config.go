package redisx

type Config struct {
	Dsn      string         `json:"dsn" yaml:"dsn"`
	Password string         `json:"password" yaml:"password"`
	DBList   map[string]int `json:"db_list" yaml:"db_list"`
}

func (c *Config) SetDefault() {
	if c.Dsn == "" {
		c.Dsn = "localhost:6379"
	}

	if len(c.DBList) == 0 {
		c.DBList = make(map[string]int)
		c.DBList["default"] = 0
	}
}
