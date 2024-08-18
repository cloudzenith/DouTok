package gofer

import "golang.org/x/sync/singleflight"

var sf *SingleFlighter

type SingleFlighter struct {
	group singleflight.Group
}

func InitSingleFlighter() {
	sf = &SingleFlighter{}
}

func (s *SingleFlighter) Do(key string, f func() (any, error)) (value any, err error, shared bool) {
	return s.group.Do(key, f)
}

func (s *SingleFlighter) DoChan(key string, f func() (any, error)) <-chan singleflight.Result {
	return s.group.DoChan(key, f)
}

func SingleFlightDo(key string, f func() (any, error)) (value any, err error, shared bool) {
	return sf.Do(key, f)
}

func SingleFlightDoChan(key string, f func() (any, error)) <-chan singleflight.Result {
	return sf.DoChan(key, f)
}

func SingleFlightForget(key string) {
	sf.group.Forget(key)
}
