package comment

type Option func(*Comment)

func WithVideoId(videoId int64) Option {
	return func(c *Comment) {
		c.VideoId = videoId
	}
}

func WithUserId(userId int64) Option {
	return func(c *Comment) {
		c.UserId = userId
	}
}

func WithParentId(parentId int64) Option {
	return func(c *Comment) {
		c.ParentId = &parentId
	}
}

func WithToUserId(toUserId int64) Option {
	return func(c *Comment) {
		c.ToUserId = &toUserId
	}
}

func WithContent(content string) Option {
	return func(c *Comment) {
		c.Content = content
	}
}
