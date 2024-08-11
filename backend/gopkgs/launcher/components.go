package launcher

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/redisx"
	"github.com/cloudzenith/DouTok/backend/gopkgs/gofer"
	"github.com/go-kratos/kratos/v2/config"
)

type ComponentsLauncher struct {
	group      *gofer.Group
	components map[string]func() error
	config     map[string]config.Value
}

func NewComponentsLauncher(config config.Config) *ComponentsLauncher {
	configMap, err := config.Value("components").Map()
	if err != nil {
		panic("get components config error: " + err.Error())
	}

	return &ComponentsLauncher{
		group: gofer.NewGroup(
			context.Background(),
			gofer.UseErrorGroup(),
		),
		components: make(map[string]func() error),
		config:     configMap,
	}
}

func (l *ComponentsLauncher) Launch() {
	launchComponent(l.config, "redis", redisx.Init)

	for _, fn := range l.components {
		l.group.Run(fn)
	}

	if err := l.group.Wait(); err != nil {
		panic("launch components error: " + err.Error())
	}
}

func launchComponent[T any](cfg map[string]config.Value, name string, initMethod func(cfg components.ConfigMap[T]) error) {
	componentConfigs, ok := cfg[name]
	if !ok {
		panic("component not found: " + name)
	}

	configs, err := componentConfigs.Map()
	if err != nil {
		panic("get component config error: " + err.Error())
	}

	if err := components.Load(configs, initMethod).Start(); err != nil {
		panic("launch component error: " + err.Error())
	}
}

func (l *ComponentsLauncher) register(name string, fn func() error) {
	l.components[name] = fn
}
