package conf

import "fmt"

type Data struct {
	Database struct {
		Source string
	}
	Redis struct {
		Source   string
		Password string
	}
	DbShardingConfig map[string]DbShardingConfig     `yaml:"db_sharding_config" json:"db_sharding_config"`
	DbShardingTables map[string]DomainShardingConfig `yaml:"db_sharding_tables" json:"db_sharding_tables"`
	Minio            struct {
		Endpoint  string
		AccessKey string
		SecretKey string
	}
}

type DbShardingConfig struct {
	Sharding       string `yaml:"sharding" json:"sharding"`
	ShardingNumber int64  `yaml:"sharding_number" json:"sharding_number"`
}

type ShardingNumberConfig struct {
	TableName      string
	DomainName     string
	BizName        string
	ShardingNumber int
}

// BizShardingFieldConfig is a map of bizName to sharding field
// e.g. {"shortVideo": ["id", "hash"]}
type BizShardingFieldConfig struct {
	Fields []string `yaml:"fields" json:"fields"`
}

// DomainShardingConfig is a map of domainName to BizShardingFieldConfig,
// used for warming up with the creation of the sharding tables
// key: domain name
type DomainShardingConfig struct {
	DomainName             string                            `yaml:"domain_name" json:"domain_name"`
	BizShardingFieldConfig map[string]BizShardingFieldConfig `yaml:"biz_sharding_field_config" json:"biz_sharding_field_config"`
}

func (config Data) GetShardingNumber(fileName, domainName, bizName string) int64 {
	key := fmt.Sprintf("%s_%s_%s", fileName, domainName, bizName)
	if num, ok := config.DbShardingConfig[key]; ok {
		return num.ShardingNumber
	}

	return 1
}
