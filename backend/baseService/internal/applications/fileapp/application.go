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
	fileId, existed, err := a.fileService.CheckFileExistedAndGetFile(ctx, request.FileContext)
	if err != nil {
		return &api.PreSignPutResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	if existed {
		return &api.PreSignPutResponse{
			Meta:   utils.GetSuccessMeta(),
			FileId: fileId,
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

func (a *FileApplication) ReportUploaded(ctx context.Context, request *api.ReportUploadedRequest) (*api.ReportUploadedResponse, error) {
	if err := a.fileService.ReportUploaded(ctx, request.FileId); err != nil {
		return &api.ReportUploadedResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.ReportUploadedResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (a *FileApplication) PreSignSlicingPut(ctx context.Context, request *api.PreSignSlicingPutRequest) (*api.PreSignSlicingPutResponse, error) {
	fileId, existed, err := a.fileService.CheckFileExistedAndGetFile(ctx, request.FileContext)
	if err != nil {
		return &api.PreSignSlicingPutResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	if existed {
		return &api.PreSignSlicingPutResponse{
			Meta:     utils.GetSuccessMeta(),
			FileId:   fileId,
			Uploaded: true,
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
	if err := a.fileService.RemoveFile(ctx, request.FileContext); err != nil {
		return &api.RemoveFileResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.RemoveFileResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (a *FileApplication) MergeFileParts(ctx context.Context, request *api.MergeFilePartsRequest) (*api.MergeFilePartsResponse, error) {
	if err := a.fileService.MergeFileParts(ctx, request.UploadId, request.FileContext); err != nil {
		return &api.MergeFilePartsResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	if err := a.fileService.ReportUploaded(ctx, request.FileContext.FileId); err != nil {
		return &api.MergeFilePartsResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.MergeFilePartsResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

var _ api.FileServiceServer = (*FileApplication)(nil)
