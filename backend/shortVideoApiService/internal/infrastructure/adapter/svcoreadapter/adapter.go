package svcoreadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/consulx"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
)

type Adapter struct {
	user       v1.UserServiceClient
	video      v1.VideoServiceClient
	collection v1.CollectionServiceClient
	comment    v1.CommentServiceClient
	favorite   v1.FavoriteServiceClient
	follow     v1.FollowServiceClient
}

func New() *Adapter {
	conn, err := consulx.GetGrpcConn(context.Background(), "discovery:///short-video-core-service")
	if err != nil {
		panic(err)
	}

	return &Adapter{
		user:       v1.NewUserServiceClient(conn),
		video:      v1.NewVideoServiceClient(conn),
		collection: v1.NewCollectionServiceClient(conn),
		comment:    v1.NewCommentServiceClient(conn),
		favorite:   v1.NewFavoriteServiceClient(conn),
		follow:     v1.NewFollowServiceClient(conn),
	}
}
