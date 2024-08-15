package repoiface

import (
	"context"
	"github.com/minio/minio-go/v7"
)

type MinioRepository interface {
	// PreSignGetUrl returns a pre-signed URL for downloading an object
	PreSignGetUrl(ctx context.Context, bucketName, objectName, fileName string, expireSeconds int64) (string, error)
	// PreSignPutUrl returns a pre-signed URL for uploading an object
	PreSignPutUrl(ctx context.Context, bucketName, objectName string, expireSeconds int64) (string, error)
	// CreateSlicingUpload creates a new multipart upload
	CreateSlicingUpload(ctx context.Context, bucketName, objectName string, options minio.PutObjectOptions) (uploadId string, err error)
	// ListSlicingFileParts lists the parts of a multipart upload
	ListSlicingFileParts(ctx context.Context, bucketName, objectName, uploadId string, partsNum int64) (minio.ListObjectPartsResult, error)
	// PreSignSlicingPutUrl returns a pre-signed URL for uploading a part of a multipart upload
	PreSignSlicingPutUrl(ctx context.Context, bucketName, objectName, uploadId string, parts int64) (string, error)
	// MergeSlices merges the parts of a multipart upload
	MergeSlices(ctx context.Context, bucketName, objectName, uploadId string, parts []minio.CompletePart) error
	// GetObjectHash returns the hash of an object
	GetObjectHash(ctx context.Context, bucketName, objectName string) (string, error)
}
