package entity

import (
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
	"time"
)

type Author struct {
	ID          int64
	Name        string
	Avatar      string
	IsFollowing int64
}

func ToAuthorEntity(author *model.User) *Author {
	if author == nil {
		return nil
	}
	return &Author{
		ID:     author.ID,
		Name:   author.Name,
		Avatar: author.Avatar,
	}
}

func (a *Author) ToPBAuthor() *v1.Author {
	return &v1.Author{
		Id:          a.ID,
		Name:        a.Name,
		Avatar:      a.Avatar,
		IsFollowing: a.IsFollowing,
	}
}

type Video struct {
	ID           int64
	Title        string
	Description  string
	VideoURL     string
	CoverURL     string
	LikeCount    int64
	CommentCount int64
	Author       *Author
	UploadTime   time.Time
	IsFavorite   int64
}

func (v *Video) ToVideoModel() *model.Video {
	return &model.Video{
		ID:           v.ID,
		UserID:       v.Author.ID,
		Title:        v.Title,
		Description:  v.Description,
		VideoURL:     v.VideoURL,
		CoverURL:     v.CoverURL,
		LikeCount:    v.LikeCount,
		CommentCount: v.CommentCount,
	}
}

func (v *Video) ToPB() *v1.Video {
	author := &v1.Author{}
	if v.Author != nil {
		author.Id = v.Author.ID
		author.Name = v.Author.Name
		author.Avatar = v.Author.Avatar
	}

	return &v1.Video{
		Id:            v.ID,
		Title:         v.Title,
		Description:   v.Description,
		PlayUrl:       v.VideoURL,
		CoverUrl:      v.CoverURL,
		IsFavorite:    v.IsFavorite,
		FavoriteCount: v.LikeCount,
		CommentCount:  v.CommentCount,
		Author:        author,
	}
}

func FromVideoModel(video *model.Video) *Video {
	if video == nil {
		return nil
	}
	return &Video{
		ID:           video.ID,
		Title:        video.Title,
		Description:  video.Description,
		VideoURL:     video.VideoURL,
		CoverURL:     video.CoverURL,
		LikeCount:    video.LikeCount,
		CommentCount: video.CommentCount,
		UploadTime:   video.CreatedAt,
		Author: &Author{
			ID: video.UserID,
		},
	}
}

func FromVideoModelList(videos []*model.Video) []*Video {
	var videoList []*Video
	for _, v := range videos {
		videoList = append(videoList, FromVideoModel(v))
	}
	return videoList
}
