package commentoptions

import v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"

type CreateCommentOption func(request *v1.CreateCommentRequest)

func CreateCommentWithVideoId(videoId int64) CreateCommentOption {
	return func(request *v1.CreateCommentRequest) {
		request.VideoId = videoId
	}
}

func CreateCommentWithUserId(userId int64) CreateCommentOption {
	return func(request *v1.CreateCommentRequest) {
		request.UserId = userId
	}
}

func CreateCommentWithContent(content string) CreateCommentOption {
	return func(request *v1.CreateCommentRequest) {
		request.Content = content
	}
}

func CreateCommentWithParentId(parentId int64) CreateCommentOption {
	return func(request *v1.CreateCommentRequest) {
		if parentId == 0 {
			return
		}

		request.ParentId = parentId
	}
}

func CreateCommentWithReplyTo(replyTo int64) CreateCommentOption {
	return func(request *v1.CreateCommentRequest) {
		if replyTo == 0 {
			return
		}

		request.ReplyUserId = replyTo
	}
}
