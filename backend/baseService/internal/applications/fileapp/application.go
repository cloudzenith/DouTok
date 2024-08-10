package fileapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/interface/fileserviceiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/utils"
)

type FileApplication struct {
	fileService fileserviceiface.FileService
}

func New(fileService fileserviceiface.FileService) *FileApplication {
	return &FileApplication{
		fileService: fileService,
	}
}

func (a *FileApplication) PreSignGet(ctx context.Context, request *api.PreSignGetRequest) (*api.PreSignGetResponse, error) {
	if err := utils.Validate(request); err != nil {
		return &api.PreSignGetResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	url, err := a.fileService.PreSignGet(ctx, request.FileContext)
	if err != nil {
		return &api.PreSignGetResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.PreSignGetResponse{
		Meta: utils.GetSuccessMeta(),
		Url:  url,
	}, nil
}

func (a *FileApplication) PreSignPut(ctx context.Context, request *api.PreSignPutRequest) (*api.PreSignPutResponse, error) {
	if err := utils.Validate(request); err != nil {
		return &api.PreSignPutResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	url, err := a.fileService.PreSignPut(ctx, request.FileContext)
	if err != nil {
		return &api.PreSignPutResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.PreSignPutResponse{
		Meta: utils.GetSuccessMeta(),
		Url:  url,
	}, nil
}

func (a *FileApplication) PreSignSlicingPut(ctx context.Context, request *api.PreSignSlicingPutRequest) (*api.PreSignSlicingPutResponse, error) {
	if err := utils.Validate(request); err != nil {
		return &api.PreSignSlicingPutResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	sf, err := a.fileService.PreSignSlicingPut(ctx, request.FileContext)
	if err != nil {
		return &api.PreSignSlicingPutResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.PreSignSlicingPutResponse{
		Meta:     utils.GetSuccessMeta(),
		Urls:     sf.UploadUrl,
		UploadId: sf.UploadId,
		Parts:    sf.TotalParts,
		FileId:   sf.File.ID,
	}, nil
}

func (a *FileApplication) GetProgressRate4SlicingPut(ctx context.Context, request *api.GetProgressRate4SlicingPutRequest) (*api.GetProgressRate4SlicingPutResponse, error) {
	if err := utils.Validate(request); err != nil {
		return &api.GetProgressRate4SlicingPutResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	result, err := a.fileService.GetProgressRate4SlicingPut(ctx, request.UploadId, request.FileContext)
	if err != nil {
		return &api.GetProgressRate4SlicingPutResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.GetProgressRate4SlicingPutResponse{
		Meta:  utils.GetSuccessMeta(),
		Parts: result,
	}, nil
}

func (a *FileApplication) RemoveFile(ctx context.Context, request *api.RemoveFileRequest) (*api.RemoveFileResponse, error) {
	if err := utils.Validate(request); err != nil {
		return &api.RemoveFileResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	if err := a.fileService.RemoveFile(ctx, request.FileContext); err != nil {
		return &api.RemoveFileResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.RemoveFileResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (a *FileApplication) ReportUploadedFileParts(ctx context.Context, request *api.ReportUploadedFilePartsRequest) (*api.ReportUploadedFilePartsResponse, error) {
	if err := utils.Validate(request); err != nil {
		return &api.ReportUploadedFilePartsResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	if err := a.fileService.ReportUploadedFileParts(ctx, request.UploadId, request.FileId, request.PartNumber); err != nil {
		return &api.ReportUploadedFilePartsResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.ReportUploadedFilePartsResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (a *FileApplication) MergeFileParts(ctx context.Context, request *api.MergeFilePartsRequest) (*api.MergeFilePartsResponse, error) {
	if err := utils.Validate(request); err != nil {
		return &api.MergeFilePartsResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	if err := a.fileService.MergeFileParts(ctx, request.UploadId, request.FileContext); err != nil {
		return &api.MergeFilePartsResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.MergeFilePartsResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

var _ api.FileServiceServer = (*FileApplication)(nil)
