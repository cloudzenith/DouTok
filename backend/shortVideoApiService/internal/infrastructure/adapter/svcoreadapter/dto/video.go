package dto

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api/svapi"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
)

func ToPBVideo(video *v1.Video) *svapi.Video {
	return &svapi.Video{
		Id:            video.GetId(),
		Title:         video.Title,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    video.IsFavorite != 0,
		Author: &svapi.VideoAuthor{
			Id:          video.Author.Id,
			Name:        video.Author.Name,
			Avatar:      video.Author.Avatar,
			IsFollowing: video.Author.IsFollowing != 0,
		},
	}
}

func ToPBVideoList(videos []*v1.Video) []*svapi.Video {
	var result []*svapi.Video
	for _, video := range videos {
		result = append(result, ToPBVideo(video))
	}
	return result
}
