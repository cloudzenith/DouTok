package videooptions

import v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"

type FeedOptions func(request *v1.FeedShortVideoRequest)

func FeedWithLatestTime(time int64) FeedOptions {
	return func(request *v1.FeedShortVideoRequest) {
		request.LatestTime = time
	}
}
