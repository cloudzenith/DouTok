package repoiface

import (
	"context"
	"github.com/minio/minio-go/v7"
)

type MinioRepository interface {
	PreSignGetUrl(ctx context.Context, bucketName, objectName, fileName string, expireSeconds int64) (string, error)
	PreSignPutUrl(ctx context.Context, bucketName, objectName string, expireSeconds int64) (string, error)
	CreateSlicingUpload(ctx context.Context, bucketName, objectName string, options minio.PutObjectOptions) (uploadId string, err error)
	ListSlicingFileParts(ctx context.Context, bucketName, objectName, uploadId string, partsNum int64) (minio.ListObjectPartsResult, error)
	PreSignSlicingPutUrl(ctx context.Context, bucketName, objectName, uploadId string, parts int64) (string, error)
	MergeSlices(ctx context.Context, bucketName, objectName, uploadId string, parts []minio.CompletePart) error
}
