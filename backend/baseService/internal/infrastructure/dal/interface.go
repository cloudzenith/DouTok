package dal

type QueryDo[T any] interface {
	FindByPage(offset int, limit int) (result []T, count int64, err error)
	Find() (result []T, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
}
