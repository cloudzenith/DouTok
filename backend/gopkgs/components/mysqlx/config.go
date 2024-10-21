package mysqlx

import "fmt"

type Config struct {
	Dialect           string `yaml:"dialect" json:"dialect"`
	Host              string `yaml:"host" json:"host"`
	Port              int    `yaml:"port" json:"port"`
	DbName            string `yaml:"db_name" json:"db_name"`
	User              string `yaml:"user" json:"user"`
	Password          string `yaml:"password" json:"password"`
	Charset           string `yaml:"charset" json:"charset"`
	ParseTime         bool   `yaml:"parse_time" json:"parse_time"`
	MaxIdle           int    `yaml:"max_idle" json:"max_idle"`
	MaxOpen           int    `yaml:"max_open" json:"max_open"`
	ConnMaxLifeTime   int    `yaml:"conn_max_life_time" json:"conn_max_life_time"`
	ConnMaxIdleTime   int    `yaml:"conn_max_idle_time" json:"conn_max_idle_time"`
	Debug             bool   `yaml:"debug" json:"debug"`
	NoLog             bool   `yaml:"no_log" json:"no_log"`
	InterpolateParams bool   `yaml:"interpolate_params" json:"interpolate_params"`
	MultiStatements   bool   `yaml:"multi_statements" json:"multi_statements"`
	Timeout           int    `yaml:"timeout" json:"timeout"`
	ReadTimeout       int    `yaml:"read_timeout" json:"read_timeout"`
	WriteTimeout      int    `yaml:"write_timeout" json:"write_timeout"`
}

func (c *Config) SetDefault() {
	if c.Dialect == "" {
		c.Dialect = "mysql"
	}

	if c.Host == "" {
		c.Host = "localhost"
	}

	if c.Port == 0 {
		c.Port = 3306
	}

	if c.Charset == "" {
		c.Charset = "utf8mb4"
	}

	if !c.ParseTime {
		c.ParseTime = true
	}

	if c.MaxIdle == 0 {
		c.MaxIdle = 10
	}

	if c.MaxOpen == 0 {
		c.MaxOpen = 100
	}
}

func (c *Config) ToDSN() string {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DbName,
		c.Charset,
		c.ParseTime,
	)

	if c.InterpolateParams {
		dsn += "&interpolateParams=true"
	}

	if c.MultiStatements {
		dsn += "&multiStatements=true"
	}

	if c.Timeout > 0 {
		dsn += fmt.Sprintf("&timeout=%ds", c.Timeout)
	}

	if c.ReadTimeout > 0 {
		dsn += fmt.Sprintf("&readTimeout=%ds", c.ReadTimeout)
	}

	if c.WriteTimeout > 0 {
		dsn += fmt.Sprintf("&writeTimeout=%ds", c.WriteTimeout)
	}

	return dsn
}
