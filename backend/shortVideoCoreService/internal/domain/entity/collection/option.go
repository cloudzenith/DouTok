package collection

type Option func(*Collection)

func WithUserId(userId int64) Option {
	return func(c *Collection) {
		c.UserId = userId
	}
}

func WithTitle(name string) Option {
	return func(c *Collection) {
		c.Title = name
	}
}

func WithDescription(description string) Option {
	return func(c *Collection) {
		c.Description = description
	}
}
