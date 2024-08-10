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
	DBShardingConfig DBShardingConfig
	Minio            struct {
		Endpoint  string
		AccessKey string
		SecretKey string
	}
}

type DBShardingConfig struct {
	ShardingNumberConfig map[string]int64
}

func (config DBShardingConfig) GetShardingNumber(domainName, bizName string) int64 {
	key := fmt.Sprintf("%s_%s", domainName, bizName)
	if num, ok := config.ShardingNumberConfig[key]; ok {
		return num
	}

	return 1
}
