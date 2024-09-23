package fileapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api/svapi"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/baseadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/errorx"
	"github.com/go-kratos/kratos/v2/log"
)

type Application struct {
	base *baseadapter.Adapter
}

func New(base *baseadapter.Adapter) *Application {
	return &Application{
		base: base,
	}
}

func (a *Application) PreSignUploadingPublicFile(ctx context.Context, request *svapi.PreSignUploadPublicFileRequest) (*svapi.PreSignUploadPublicFileResponse, error) {
	resp, err := a.base.PreSign4PublicUpload(
		ctx,
		request.Hash,
		request.FileType,
		request.FileType,
		request.Size,
		3600,
	)
	if err != nil {
		log.Context(ctx).Errorf("failed to presign: %v", err)
		return nil, errorx.New(1, "failed to presign")
	}

	return &svapi.PreSignUploadPublicFileResponse{
		Url:    resp.Url,
		FileId: resp.FileId,
	}, nil
}

func (a *Application) ReportPublicFileUploaded(ctx context.Context, request *svapi.ReportPublicFileUploadedRequest) (*svapi.ReportPublicFileUploadedResponse, error) {
	_, err := a.base.ReportPublicUploaded(ctx, request.FileId)
	if err != nil {
		log.Context(ctx).Errorf("failed to report uploaded: %v", err)
		return nil, errorx.New(1, "failed to report uploaded")
	}

	info, err := a.base.GetFileInfoById(ctx, request.FileId)
	if err != nil {
		log.Context(ctx).Errorf("failed to get file info: %v", err)
		return nil, errorx.New(1, "failed to get file info")
	}

	return &svapi.ReportPublicFileUploadedResponse{
		ObjectName: info.ObjectName,
	}, nil
}

var _ svapi.FileServiceHTTPServer = (*Application)(nil)
