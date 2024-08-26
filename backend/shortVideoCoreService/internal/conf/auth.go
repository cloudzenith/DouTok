package conf

type Auth struct {
	JWT struct {
		AccessSecret string `yaml:"access_secret" json:"access_secret"`
		AccessExpire int64  `yaml:"access_expire" json:"access_expire"`
	} `yaml:"jwt" json:"jwt"`
}
