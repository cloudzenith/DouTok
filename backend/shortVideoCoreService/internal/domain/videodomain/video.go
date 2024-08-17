package videodomain

import (
	"context"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	domain_dto "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/dto"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/videodata"
	service_dto "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/dto"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/userdomain"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/db"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/pkg/auth"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/pkg/utils"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type VideoUseCase struct {
	config    *conf.Config
	videoRepo videodata.IVideoRepo
	userRepo  userdomain.UserRepo
	snowflake *utils.SnowflakeNode
	dbClient  *db.DBClient
	log       *log.Helper
}

func NewVideoUseCase(
	config *conf.Config,
	snowflake *utils.SnowflakeNode,
	userRepo userdomain.UserRepo,
	videoRepo videodata.IVideoRepo,
	dbClient *db.DBClient,
	logger log.Logger,
) *VideoUseCase {
	return &VideoUseCase{
		config:    config,
		videoRepo: videoRepo,
		userRepo:  userRepo,
		snowflake: snowflake,
		dbClient:  dbClient,
		log:       log.NewHelper(logger),
	}
}

func (uc *VideoUseCase) FeedShortVideo(ctx context.Context, request *service_dto.FeedShortVideoRequest) (*service_dto.FeedShortVideoResponse, error) {
	latestTime := time.Now().UTC().Unix()
	if request.LatestTime > 0 {
		latestTime = request.LatestTime
	}

	resp, err := uc.videoRepo.GetVideoFeed(ctx, &domain_dto.GetVideoFeedRequest{
		UserId:     request.UserId,
		LatestTime: latestTime,
		Num:        request.FeedNum,
	})
	if err != nil {
		return nil, err
	}
	user, err := uc.userRepo.FindByID(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	videos := entity.FromVideoModelList(resp.Videos)
	author := entity.ToAuthorEntity(user)
	for _, video := range videos {
		video.Author = author
	}

	return &service_dto.FeedShortVideoResponse{
		Videos: videos,
	}, nil
}

func (uc *VideoUseCase) PublishVideo(ctx context.Context, in *service_dto.PublishVideoRequest) (int64, error) {
	userId, err := auth.GetLoginUser(ctx)
	if err != nil {
		return 0, fmt.Errorf("get login user failed: %v", err)
	}
	video := model.Video{
		ID:          uc.snowflake.GetSnowflakeId(),
		UserID:      userId,
		Title:       in.Title,
		Description: in.Description,
		VideoURL:    in.VideoURL,
		CoverURL:    in.CoverURL,
	}
	err = uc.videoRepo.Save(ctx, &video)
	if err != nil {
		return 0, err
	}
	return video.ID, nil
}

func (uc *VideoUseCase) GetVideoById(ctx context.Context, videoId int64) (*entity.Video, error) {
	video, err := uc.videoRepo.FindByID(ctx, videoId)
	if err != nil {
		return nil, err
	}
	authorModel, err := uc.userRepo.FindByID(ctx, video.UserID)
	if err != nil {
		return nil, err
	}
	videoEntity := entity.FromVideoModel(video)
	videoEntity.Author = entity.ToAuthorEntity(authorModel)
	return videoEntity, nil
}

func (uc *VideoUseCase) ListPublishedVideo(ctx context.Context, request *service_dto.ListPublishedVideoRequest) (*service_dto.ListPublishedVideoResponse, error) {
	latestTime := time.Now().UTC().Unix()
	if request.LatestTime > 0 {
		latestTime = request.LatestTime
	}

	user, err := uc.userRepo.FindByID(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	resp, err := uc.videoRepo.GetVideoList(ctx, &domain_dto.GetVideoListRequest{
		UserId:            request.UserId,
		LatestTime:        latestTime,
		PaginationRequest: request.Pagination,
	})
	if err != nil {
		return nil, err
	}

	videos := entity.FromVideoModelList(resp.Videos)
	author := entity.ToAuthorEntity(user)
	for _, video := range videos {
		video.Author = author
	}

	return &service_dto.ListPublishedVideoResponse{
		Videos:     videos,
		Pagination: resp.PaginationResponse,
	}, nil
}
