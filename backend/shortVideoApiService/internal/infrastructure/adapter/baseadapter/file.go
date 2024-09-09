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

type PreSign4UploadResp struct {
	Url    string
	FileId int64
}

type ReportUploadedResp struct {
	Url string
}

func (a *Adapter) PreSign4Upload(ctx context.Context, hash, fileType, fileName string, size, expireSeconds int64) (*PreSign4UploadResp, error) {
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
	return respcheck.CheckT[*PreSign4UploadResp, *api.Metadata](
		resp, err,
		func() *PreSign4UploadResp {
			return &PreSign4UploadResp{
				Url:    resp.GetUrl(),
				FileId: resp.GetFileId(),
			}
		},
	)
}

func (a *Adapter) ReportUploaded(ctx context.Context, fileId int64) (*ReportUploadedResp, error) {
	req := &api.ReportUploadedRequest{
		FileContext: &api.FileContext{
			BizName:       BizName,
			Domain:        DomainName,
			FileId:        fileId,
			ExpireSeconds: 7200,
		},
	}
	resp, err := a.file.ReportUploaded(ctx, req)
	return respcheck.CheckT[*ReportUploadedResp, *api.Metadata](
		resp, err,
		func() *ReportUploadedResp {
			return &ReportUploadedResp{
				Url: resp.Url,
			}
		},
	)
}
