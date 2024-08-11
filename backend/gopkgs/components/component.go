package components

import "github.com/go-kratos/kratos/v2/config"

type ConfigMap[T any] map[string]T

type Component[T any] struct {
	err         error
	configValue map[string]config.Value
	cfg         ConfigMap[T]
	initMethod  func(cfg ConfigMap[T]) error
}

func Load[T any](configValue map[string]config.Value, initMethod func(cfg ConfigMap[T]) error) *Component[T] {
	c := &Component[T]{
		initMethod:  initMethod,
		configValue: configValue,
	}

	c.cfg = make(ConfigMap[T])
	for key, value := range configValue {
		var v interface{}
		if err := value.Scan(v); err != nil {
			panic("scan component config error: " + err.Error())
		}
		c.cfg[key] = v.(T)
	}

	return c
}

func (s *Component[T]) Start() error {
	if s.err != nil {
		return s.err
	}

	return s.initMethod(s.cfg)
}

func (s *Component[T]) GetConfig() ConfigMap[T] {
	return s.cfg
}
