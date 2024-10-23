package videoservice

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api/svapi"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/interface/videoserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type VideoService struct {
	core *svcoreadapter.Adapter
}

func New(core *svcoreadapter.Adapter) *VideoService {
	return &VideoService{
		core: core,
	}
}

func (s *VideoService) AssembleVideo(ctx context.Context, userId int64, data []*v1.Video) ([]*svapi.Video, error) {
	videoList, _ := s.AssembleVideoList(ctx, userId, data)
	s.AssembleAuthorInfo(ctx, videoList)
	s.AssembleUserIsFollowing(ctx, videoList, userId)
	s.AssembleVideoCountInfo(ctx, videoList)
	return videoList, nil
}

func (s *VideoService) AssembleAuthorInfo(ctx context.Context, data []*svapi.Video) {
	var userIdList []int64
	for _, video := range data {
		if video.Author != nil {
			userIdList = append(userIdList, video.Author.Id)
		}
	}

	userList, err := s.core.GetUserInfoByIdList(ctx, userIdList)
	if err != nil {
		log.Context(ctx).Warnf("failed to get user info: %v", err)
	}

	userMap := make(map[int64]*v1.User)
	for _, user := range userList {
		userMap[user.Id] = user
	}

	for _, video := range data {
		if video.Author == nil {
			continue
		}

		author, ok := userMap[video.Author.Id]
		if !ok {
			continue
		}
		video.Author.Name = author.Name
		video.Author.Avatar = author.Avatar
	}
}

func (s *VideoService) AssembleVideoList(ctx context.Context, userId int64, data []*v1.Video) ([]*svapi.Video, error) {
	var result []*svapi.Video

	var videoIdList []int64
	for _, video := range data {
		videoIdList = append(videoIdList, video.GetId())
	}

	isFavoriteMap, err := s.core.IsUserFavoriteVideo(ctx, userId, videoIdList)
	if err != nil {
		// 弱依赖
		log.Context(ctx).Warnf("failed to check favorite video: %v", err)
	}

	for _, video := range data {
		isFavorite, ok := isFavoriteMap[video.GetId()]

		result = append(result, &svapi.Video{
			Id:            video.GetId(),
			Title:         video.Title,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite && ok,
			Author: &svapi.VideoAuthor{
				Id:          video.Author.Id,
				Name:        video.Author.Name,
				Avatar:      video.Author.Avatar,
				IsFollowing: video.Author.IsFollowing != 0,
			},
		})
	}

	return result, nil
}

func (s *VideoService) AssembleUserIsFollowing(ctx context.Context, list []*svapi.Video, userId int64) {
	var targetUserId []int64
	var targetVideoId []int64
	for _, video := range list {
		targetUserId = append(targetUserId, video.GetAuthor().GetId())
		targetVideoId = append(targetVideoId, video.GetId())
	}

	isFollowingMap, err := s.core.IsFollowing(ctx, userId, targetUserId)
	if err != nil {
		log.Context(ctx).Errorf("failed to check is following: %v", err)
	}

	isCollectedMap, err := s.core.IsCollected(ctx, userId, targetVideoId)
	if err != nil {
		log.Context(ctx).Errorf("failed to check is collected: %v", err)
	}

	isFavoriteMap, err := s.core.IsUserFavoriteVideo(ctx, userId, targetVideoId)
	if err != nil {
		log.Context(ctx).Errorf("failed to check is favorite: %v", err)
	}

	for _, video := range list {
		author := video.GetAuthor()
		author.IsFollowing = isFollowingMap[author.GetId()]
		video.IsCollected = isCollectedMap[video.GetId()]
		video.IsFavorite = isFavoriteMap[video.GetId()]
	}
}

func (s *VideoService) AssembleVideoCountInfo(ctx context.Context, list []*svapi.Video) {
	var videoIdList []int64
	for _, video := range list {
		videoIdList = append(videoIdList, video.GetId())
	}

	commentCountMap, err := s.core.CountComments4Video(ctx, videoIdList)
	if err != nil {
		log.Context(ctx).Errorf("failed to count comments: %v", err)
	}

	favoriteCountMap, err := s.core.CountFavorite4Video(ctx, videoIdList)
	if err != nil {
		log.Context(ctx).Errorf("failed to count favorite: %v", err)
	}

	collectedCountMap, err := s.core.CountCollected4Video(ctx, videoIdList)
	if err != nil {
		log.Context(ctx).Errorf("failed to count collected: %v", err)
	}

	for _, video := range list {
		video.CommentCount = commentCountMap[video.GetId()]
		video.FavoriteCount = favoriteCountMap[video.GetId()]
		video.CollectedCount = collectedCountMap[video.GetId()]
	}
}

var _ videoserviceiface.VideoService = &VideoService{}
