package dto

import (
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity"
	"time"
)

func ToPBVideo(v *entity.Video) *v1.Video {
	return &v1.Video{
		Id:            v.ID,
		Title:         v.Title,
		Author:        v.Author.ToPBAuthor(),
		PlayUrl:       v.VideoURL,
		CoverUrl:      v.CoverURL,
		FavoriteCount: v.LikeCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    v.IsFavorite,
		UploadTime:    v.UploadTime.Format(time.DateTime),
	}
}

func ToPBVideoList(videos []*entity.Video) []*v1.Video {
	var pbVideos []*v1.Video
	for _, v := range videos {
		pbVideos = append(pbVideos, ToPBVideo(v))
	}
	return pbVideos
}
