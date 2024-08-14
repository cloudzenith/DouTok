package fileserviceiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/slicingfile"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/service/fileservice"
)

//go:generate mockgen -source=interface.go -destination=interface_mock.go -package=fileserviceiface FileService
type FileService interface {
	PreSignGet(ctx context.Context, fileCtx *api.FileContext) (string, error)
	PreSignPut(ctx context.Context, fileCtx *api.FileContext) (string, error)
	CheckFileExistedAndGetFile(ctx context.Context, fileCtx *api.FileContext) (int64, bool, error)
	ReportUploaded(ctx context.Context, fileId int64) error
	PreSignSlicingPut(ctx context.Context, fileCtx *api.FileContext) (*slicingfile.SlicingFile, error)
	GetProgressRate4SlicingPut(ctx context.Context, uploadId string, fileCtx *api.FileContext) (map[string]bool, error)
	MergeFileParts(ctx context.Context, uploadId string, fileCtx *api.FileContext) error
	RemoveFile(ctx context.Context, fileCtx *api.FileContext) error
}

var _ FileService = (*fileservice.FileService)(nil)
