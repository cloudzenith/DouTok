package fileserviceiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/file"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/slicingfile"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/service/fileservice"
)

//go:generate mockgen -source=interface.go -destination=interface_mock.go -package=fileserviceiface FileService
type FileService interface {
	// PreSignGet returns the pre-signed URL for downloading the file
	PreSignGet(ctx context.Context, fileCtx *api.FileContext) (string, error)
	// PreSignPut returns the pre-signed URL for uploading the file, and the file ID if the file already exists
	PreSignPut(ctx context.Context, fileCtx *api.FileContext) (string, int64, error)
	// CheckFileExistedAndGetFile checks if the file exists and returns the file ID if it does
	CheckFileExistedAndGetFile(ctx context.Context, fileCtx *api.FileContext) (int64, bool, error)
	// ReportUploaded reports that the file has been uploaded
	ReportUploaded(ctx context.Context, fileCtx *api.FileContext) error
	// PreSignSlicingPut returns the pre-signed URL for uploading the file in slices
	PreSignSlicingPut(ctx context.Context, fileCtx *api.FileContext) (*slicingfile.SlicingFile, error)
	// GetProgressRate4SlicingPut returns the progress rate of the file upload in slices
	GetProgressRate4SlicingPut(ctx context.Context, uploadId string, fileCtx *api.FileContext) (map[string]bool, error)
	// MergeFileParts merges the file parts uploaded in slices
	MergeFileParts(ctx context.Context, uploadId string, fileCtx *api.FileContext) error
	// RemoveFile removes the file
	RemoveFile(ctx context.Context, fileCtx *api.FileContext) error
	GetInfoById(ctx context.Context, domainName, bizName string, fileId int64) (*file.File, error)
}

var _ FileService = (*fileservice.FileService)(nil)
