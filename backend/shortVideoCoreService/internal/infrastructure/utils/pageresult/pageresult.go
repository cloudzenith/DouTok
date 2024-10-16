package pageresult

type R[T any] struct {
	Data  []T
	Count int64
}

func New[T any](data []T, count int64) *R[T] {
	return &R[T]{
		Data:  data,
		Count: count,
	}
}
