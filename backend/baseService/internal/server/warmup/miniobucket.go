package warmup

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/minio/minio-go/v7"
)

func CheckAndCreateMinioBucket(core *minio.Core, shardingConfigs map[string]conf.DomainShardingConfig) {
	needCheckBucket := make(map[string]bool)
	for _, domainConfig := range shardingConfigs {
		if _, ok := needCheckBucket[domainConfig.DomainName]; !ok {
			needCheckBucket[domainConfig.DomainName] = true
		}
	}

	for bucketName := range needCheckBucket {
		checkAndCreate(core, bucketName)
	}
}

func checkAndCreate(core *minio.Core, bucketName string) {
	log.Infof("check and create bucket %s", bucketName)
	exist, err := core.BucketExists(context.Background(), bucketName)

	if exist && err == nil {
		log.Infof("bucket %s already exists", bucketName)
		return
	}

	if err := core.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{}); err != nil {
		log.Errorf("create bucket %s failed: %v", bucketName, err)
		panic(err)
	}

	log.Infof("create bucket %s success", bucketName)
}
