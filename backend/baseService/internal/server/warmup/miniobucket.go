package warmup

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/cors"
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
	if err := core.SetBucketCors(context.Background(), bucketName, &cors.Config{
		CORSRules: []cors.Rule{
			{
				AllowedOrigin: []string{"*"},
				AllowedMethod: []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
				AllowedHeader: []string{"*"},
				ExposeHeader:  []string{"ETag"},
			},
		},
	}); err != nil {
		log.Errorf("set bucket %s cors failed: %v", bucketName, err)
	}

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
