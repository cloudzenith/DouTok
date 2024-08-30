package conf

type Config struct {
	App        App        `yaml:"app" json:"app"`
	Server     Server     `yaml:"server" json:"server"`
	Auth       Auth       `yaml:"auth" json:"auth"`
	Components Components `yaml:"components" json:"components"`
}
