package warmup

import (
	"context"
	"fmt"
	minio2 "github.com/cloudzenith/DouTok/backend/gopkgs/components/miniox"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/minio/minio-go/v7"
	"os/exec"
)

func InitMinioPublicDirectory() {
	entrypoint := "http://minio:9000"
	username := "root"
	password := "rootroot"
	bucketNameList := []string{"shortvideo"}
	if output, err := exec.Command("./mc", "config", "host", "add", "minio", entrypoint, username, password).CombinedOutput(); err != nil {
		log.Errorf("mc config host add err: %v, output: %s", err, output)
		panic(err)
	}
	for _, bucketName := range bucketNameList {
		if output, err := exec.Command("./mc", "anonymous", "set", "public", "minio/"+bucketName+"/public").CombinedOutput(); err != nil {
			log.Errorf("mc anonymous set err: %v, output: %s", err, output)
			panic(err)
		}
	}
}

func InitMinioPublicDirectoryV2() {
	// 存储桶列表，可以从配置读取，调试目的，这里先写死
	bucketNameList := []string{"shortvideo"}

	// 获取 MinIO 客户端
	minioClient := minio2.GetClient(context.Background())

	// 设置存储桶策略，允许匿名用户读取
	policyStatement := `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Principal": "*",
				"Action": [
					"s3:*"
				],
				"Resource": [
					"arn:aws:s3:::%s/public/*"
				]
			}
		]
	}`

	for _, bucketName := range bucketNameList {
		// 创建存储桶
		err := minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			exist, _ := minioClient.BucketExists(context.Background(), bucketName)
			if exist {
				log.Infof("Bucket %s already exists.", bucketName)
			} else {
				log.Errorf("Error creating bucket %s: %v", bucketName, err)
			}
		}

		// 设置存储桶策略
		policy := fmt.Sprintf(policyStatement, bucketName)
		if err := minioClient.SetBucketPolicy(context.Background(), bucketName, policy); err != nil {
			log.Errorf("Error setting bucket policy for bucket %s: %v", bucketName, err)
			return
		}
	}
}
