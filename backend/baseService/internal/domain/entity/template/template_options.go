package template

type Option func(*Template)

func WithID(id int64) Option {
	return func(t *Template) {
		t.ID = id
	}
}

func WithTitle(title string) Option {
	return func(t *Template) {
		t.Title = title
	}
}

func WithContent(content string) Option {
	return func(t *Template) {
		t.Content = content
	}
}

func WithIsDelete(isDelete bool) Option {
	return func(t *Template) {
		t.IsDelete = isDelete
	}
}
