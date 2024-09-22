package baseadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/respcheck"
)

const (
	DomainName = "shortvideo"
	BizName    = "short_video"
	Public     = "public"
)

type PreSign4UploadResp struct {
	Url    string
	FileId int64
}

type ReportUploadedResp struct {
	Url string
}

func (a *Adapter) preSign4Upload(ctx context.Context, domainName, bizName, hash, fileType, fileName string, size, expireSeconds int64) (*PreSign4UploadResp, error) {
	req := &api.PreSignPutRequest{
		FileContext: &api.FileContext{
			Domain:        domainName,
			BizName:       bizName,
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

func (a *Adapter) PreSign4Upload(ctx context.Context, hash, fileType, fileName string, size, expireSeconds int64) (*PreSign4UploadResp, error) {
	return a.preSign4Upload(ctx, DomainName, BizName, hash, fileType, fileName, size, expireSeconds)
}

func (a *Adapter) PreSign4PublicUpload(ctx context.Context, hash, fileType, fileName string, size, expireSeconds int64) (*PreSign4UploadResp, error) {
	return a.preSign4Upload(ctx, DomainName, Public, hash, fileType, fileName, size, expireSeconds)
}

func (a *Adapter) reportUploaded(ctx context.Context, domainName, bizName string, fileId int64) (*ReportUploadedResp, error) {
	req := &api.ReportUploadedRequest{
		FileContext: &api.FileContext{
			BizName:       bizName,
			Domain:        domainName,
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

func (a *Adapter) ReportUploaded(ctx context.Context, fileId int64) (*ReportUploadedResp, error) {
	return a.reportUploaded(ctx, DomainName, BizName, fileId)
}

func (a *Adapter) ReportPublicUploaded(ctx context.Context, fileId int64) (*ReportUploadedResp, error) {
	return a.reportUploaded(ctx, DomainName, Public, fileId)
}

func (a *Adapter) GetFileInfoById(ctx context.Context, fileId int64) (*api.GetFileInfoByIdResponse, error) {
	req := &api.GetFileInfoByIdRequest{
		DomainName: DomainName,
		BizName:    Public,
		FileId:     fileId,
	}
	resp, err := a.file.GetFileInfoById(ctx, req)
	return respcheck.CheckT[*api.GetFileInfoByIdResponse, *api.Metadata](
		resp, err,
		func() *api.GetFileInfoByIdResponse {
			return resp
		},
	)
}
