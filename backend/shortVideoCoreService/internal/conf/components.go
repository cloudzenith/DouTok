package conf

type Components struct {
	MySQL  MySQL  `yaml:"mysql" json:"mysql"`
	Redis  Redis  `yaml:"redis" json:"redis"`
	Consul Consul `yaml:"consul" json:"consul"`
}

type MySQL struct {
	Default struct {
		Host     string `yaml:"host" json:"host"`
		Port     int    `yaml:"port" json:"port"`
		DBName   string `yaml:"db_name" json:"db_name"`
		User     string `yaml:"user" json:"user"`
		Password string `yaml:"password" json:"password"`
	} `yaml:"default" json:"default"`
}

type Redis struct {
	Default struct {
		DSN      string `yaml:"dsn" json:"dsn"`
		Password string `yaml:"password" json:"password"`
	} `yaml:"default" json:"default"`
}

type Consul struct {
	Default struct {
		Address string `yaml:"address" json:"address"`
	}
}
