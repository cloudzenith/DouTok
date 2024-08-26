package conf

type Components struct {
	MySQL struct {
		Default struct {
			Source string `yaml:"source" json:"source"`
		} `yaml:"default" json:"default"`
	} `yaml:"mysql" json:"mysql"`
	Redis struct {
		Default struct {
			DSN      string `yaml:"dsn" json:"dsn"`
			Password string `yaml:"password" json:"password"`
		} `yaml:"default" json:"default"`
	} `yaml:"redis" json:"redis"`
	Consul struct {
		Default struct {
			Address string `yaml:"address" json:"address"`
		} `yaml:"default" json:"default"`
	} `yaml:"consul" json:"consul"`
	BaseService struct {
		Address string `yaml:"address" json:"address"`
	} `yaml:"base_service" json:"base_service"`
}
