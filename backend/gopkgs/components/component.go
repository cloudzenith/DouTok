package components

import "github.com/go-kratos/kratos/v2/config"

type ConfigMap[T any] map[string]T

type Component[T any] struct {
	err         error
	configValue map[string]config.Value
	cfg         ConfigMap[T]
	initMethod  func(cfg ConfigMap[T]) (func() error, error)
}

func Load[T any](configValue map[string]config.Value, initMethod func(cfg ConfigMap[*T]) (func() error, error)) (t *T, components *Component[*T]) {
	c := &Component[*T]{
		initMethod:  initMethod,
		configValue: configValue,
	}

	c.cfg = make(ConfigMap[*T])
	for key, value := range configValue {
		t = new(T)
		if err := value.Scan(t); err != nil {
			panic("scan component config error: " + err.Error())
		}
		c.cfg[key] = t
	}

	return t, c
}

func (s *Component[T]) Start() error {
	if s.err != nil {
		return s.err
	}

	healthCheckMethod, err := s.initMethod(s.cfg)
	if err != nil {
		return err
	}

	return healthCheckMethod()
}

func (s *Component[T]) GetConfig() ConfigMap[T] {
	return s.cfg
}
