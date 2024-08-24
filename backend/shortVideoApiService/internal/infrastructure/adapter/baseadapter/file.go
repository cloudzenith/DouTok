package baseadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/respcheck"
)

const (
	DomainName = "shortvideo"
	BizName    = "short_video"
)

type PreSign4UploadVideoResp struct {
	Url    string
	FileId int64
}

func (a *Adapter) PreSign4UploadVideo(ctx context.Context, hash, fileType, fileName string, size, expireSeconds int64) (*PreSign4UploadVideoResp, error) {
	req := &api.PreSignPutRequest{
		FileContext: &api.FileContext{
			Domain:        DomainName,
			BizName:       BizName,
			Hash:          hash,
			FileType:      fileType,
			Size:          size,
			ExpireSeconds: expireSeconds,
			Filename:      fileName,
		},
	}
	resp, err := a.file.PreSignPut(ctx, req)
	return respcheck.CheckT[*PreSign4UploadVideoResp, *api.Metadata](
		resp, err,
		func() *PreSign4UploadVideoResp {
			return &PreSign4UploadVideoResp{
				Url:    resp.GetUrl(),
				FileId: resp.GetFileId(),
			}
		},
	)
}

func (a *Adapter) ReportUploaded(ctx context.Context, fileId int64) error {
	req := &api.ReportUploadedRequest{
		FileContext: &api.FileContext{
			FileId: fileId,
		},
	}
	resp, err := a.file.ReportUploaded(ctx, req)
	return respcheck.Check[*api.Metadata](resp, err)
}
