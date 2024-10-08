package commentappprovider

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/commentapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/service/commentservice"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/repositories/commentrepo"
)

func InitCommentApplication() *commentapp.Application {
	commentRepo := commentrepo.New()
	commentService := commentservice.New(commentRepo)
	commentApp := commentapp.New(commentService)
	return commentApp
}
