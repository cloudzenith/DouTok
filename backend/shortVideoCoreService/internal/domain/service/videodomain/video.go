package videodomain

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/userdata"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/videodata"
	service_dto "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/dto"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/query"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type VideoUseCase struct {
	config    *conf.Config
	videoRepo videodata.IVideoRepo
	userRepo  userdata.IUserRepo
}

func NewVideoUseCase(
	config *conf.Config,
	userRepo userdata.IUserRepo,
	videoRepo videodata.IVideoRepo,
) *VideoUseCase {
	return &VideoUseCase{
		config:    config,
		videoRepo: videoRepo,
		userRepo:  userRepo,
	}
}

func (uc *VideoUseCase) FeedShortVideo(ctx context.Context, request *service_dto.FeedShortVideoRequest) (*service_dto.FeedShortVideoResponse, error) {
	latestTime := time.Now().UTC().Unix()
	if request.LatestTime > 0 {
		latestTime = request.LatestTime
	}

	videos, err := uc.videoRepo.GetVideoFeed(ctx, query.Q, request.UserId, latestTime, request.FeedNum)
	if err != nil {
		return nil, err
	}

	// 去重并查询用户
	uniqueUserIds := make(map[int64]struct{})
	for _, video := range videos {
		uniqueUserIds[video.UserID] = struct{}{}
	}
	userIds := make([]int64, 0, len(uniqueUserIds))
	for id := range uniqueUserIds {
		userIds = append(userIds, id)
	}
	users, err := uc.userRepo.FindByIds(ctx, query.Q, userIds)
	if err != nil {
		return nil, err
	}

	// 构建用户映射
	userMap := make(map[int64]*model.User, len(users))
	for _, user := range users {
		userMap[user.ID] = user
	}

	// 构建视频列表
	videoList := make([]*entity.Video, 0, len(videos))
	for _, videoModel := range videos {
		videoEntity := entity.FromVideoModel(videoModel)
		authorModel, ok := userMap[videoModel.UserID]
		if !ok {
			log.Warnf("user not found: %d", videoModel.UserID)
		}
		videoEntity.Author = entity.ToAuthorEntity(authorModel)
		videoList = append(videoList, videoEntity)
	}

	return &service_dto.FeedShortVideoResponse{
		Videos: videoList,
	}, nil
}

func (uc *VideoUseCase) PublishVideo(ctx context.Context, in *service_dto.PublishVideoRequest) (int64, error) {
	video := model.Video{
		ID:          utils.GetSnowflakeId(),
		UserID:      in.UserId,
		Title:       in.Title,
		Description: in.Description,
		VideoURL:    in.VideoURL,
		CoverURL:    in.CoverURL,
	}
	err := uc.videoRepo.Save(ctx, query.Q, &video)
	if err != nil {
		return 0, err
	}
	return video.ID, nil
}

func (uc *VideoUseCase) GetVideoById(ctx context.Context, videoId int64) (*entity.Video, error) {
	video, err := uc.videoRepo.FindByID(ctx, query.Q, videoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("video not found: %d", videoId)
	}
	if err != nil {
		return nil, err
	}
	authorModel, err := uc.userRepo.FindByID(ctx, query.Q, video.UserID)
	if err != nil {
		return nil, err
	}
	videoEntity := entity.FromVideoModel(video)
	videoEntity.Author = entity.ToAuthorEntity(authorModel)
	return videoEntity, nil
}

func (uc *VideoUseCase) GetVideoByIdList(ctx context.Context, videoIdList []int64) ([]*entity.Video, error) {
	data, err := uc.videoRepo.FindByIdList(ctx, videoIdList)
	if err != nil {
		log.Context(ctx).Errorf("GetVideoByIdList error: %v", err)
		return nil, err
	}

	var result []*entity.Video
	for _, item := range data {
		result = append(result, entity.FromVideoModel(item))
	}

	return result, nil
}

func (uc *VideoUseCase) ListPublishedVideo(ctx context.Context, request *service_dto.ListPublishedVideoRequest) (*service_dto.ListPublishedVideoResponse, error) {
	latestTime := time.Now().UTC().Unix()
	if request.LatestTime > 0 {
		latestTime = request.LatestTime
	}

	user, err := uc.userRepo.FindByID(ctx, query.Q, request.UserId)
	if err != nil {
		return nil, err
	}

	videos, pageResp, err := uc.videoRepo.GetVideoList(ctx, query.Q, request.UserId, latestTime, request.Pagination)
	if err != nil {
		return nil, err
	}

	videoEntityList := entity.FromVideoModelList(videos)
	author := entity.ToAuthorEntity(user)
	for _, video := range videoEntityList {
		video.Author = author
	}

	return &service_dto.ListPublishedVideoResponse{
		Videos:     videoEntityList,
		Pagination: pageResp,
	}, nil
}
