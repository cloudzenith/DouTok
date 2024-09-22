package file

type Option func(*File)

func WithID(id int64) Option {
	return func(f *File) {
		f.ID = id
	}
}

func WithDomain(domain string) Option {
	return func(f *File) {
		f.DomainName = domain
	}
}

func WithBizName(bizName string) Option {
	return func(f *File) {
		f.BizName = bizName
	}
}

func WithHash(hash string) Option {
	return func(f *File) {
		f.Hash = hash
	}
}

func WithFileType(fileType string) Option {
	return func(f *File) {
		f.FileType = fileType
	}
}

func WithSize(size int64) Option {
	return func(f *File) {
		f.FileSize = size
	}
}
