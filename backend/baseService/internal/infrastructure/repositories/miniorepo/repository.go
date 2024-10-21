package miniorepo

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/miniox"
	"github.com/minio/minio-go/v7"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type PersistRepository struct {
	core *minio.Core
}

func New() *PersistRepository {
	return &PersistRepository{
		core: miniox.GetClient(context.Background()),
	}
}

func (r *PersistRepository) PreSignGetUrl(ctx context.Context, bucketName, objectName, fileName string, expireSeconds int64) (string, error) {
	reqParams := make(url.Values)

	if fileName != "" {
		reqParams.Set("response-content-disposition", "attachment; filename="+fileName)
	}

	signedUrl, err := r.core.PresignedGetObject(ctx, bucketName, objectName, time.Duration(expireSeconds*1000000000), reqParams)
	if err != nil {
		return "", err
	}

	return signedUrl.String(), nil
}

func (r *PersistRepository) PreSignPutUrl(ctx context.Context, bucketName, objectName string, expireSeconds int64) (string, error) {
	signedUrl, err := r.core.PresignedPutObject(ctx, bucketName, objectName, time.Duration(expireSeconds*1000000000))
	if err != nil {
		return "", err
	}

	return signedUrl.String(), nil
}

func (r *PersistRepository) CreateSlicingUpload(ctx context.Context, bucketName, objectName string, options minio.PutObjectOptions) (uploadId string, err error) {
	return r.core.NewMultipartUpload(ctx, bucketName, objectName, options)
}

func (r *PersistRepository) ListSlicingFileParts(ctx context.Context, bucketName, objectName, uploadId string, partsNum int64) (minio.ListObjectPartsResult, error) {
	var nextPartNumberMarker int
	return r.core.ListObjectParts(ctx, bucketName, objectName, uploadId, nextPartNumberMarker, int(partsNum)+1)
}

func (r *PersistRepository) PreSignSlicingPutUrl(ctx context.Context, bucketName, objectName, uploadId string, parts int64) (string, error) {
	params := url.Values{
		"uploadId":   {uploadId},
		"partNumber": {strconv.FormatInt(parts, 10)},
	}

	signedUrl, err := r.core.Presign(ctx, http.MethodPut, bucketName, objectName, time.Hour, params)
	if err != nil {
		return "", err
	}

	return signedUrl.String(), nil
}

func (r *PersistRepository) MergeSlices(ctx context.Context, bucketName, objectName, uploadId string, parts []minio.CompletePart) error {
	_, err := r.core.CompleteMultipartUpload(ctx, bucketName, objectName, uploadId, parts, minio.PutObjectOptions{})
	return err
}

func (r *PersistRepository) GetObjectHash(ctx context.Context, bucketName, objectName string) (string, error) {
	stats, err := r.core.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		return "", err
	}

	return strings.ToUpper(stats.ETag), nil
}
