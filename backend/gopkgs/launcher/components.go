package launcher

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/consulx"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/etcdx"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/miniox"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/mysqlx"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/redisx"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/rmqconsumerx"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/rmqproducerx"
	"github.com/cloudzenith/DouTok/backend/gopkgs/gofer"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
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

func launchWrapper(cfg config.Value, componentsName string) {
	if componentsName == "mysql" {
		launchComponent(cfg, mysqlx.Init)
		return
	}

	if componentsName == "redis" {
		launchComponent(cfg, redisx.Init)
		return
	}

	if componentsName == "minio" {
		launchComponent(cfg, miniox.Init)
		return
	}

	if componentsName == "etcd" {
		launchComponent(cfg, etcdx.Init)
		return
	}

	if componentsName == "consul" {
		launchComponent(cfg, consulx.Init)
		return
	}

	if componentsName == "rmqconsumer" {
		launchComponent(cfg, rmqconsumerx.Init)
		return
	}

	if componentsName == "rmqproducer" {
		launchComponent(cfg, rmqproducerx.Init)
		return
	}

	panic("unknown components name: " + componentsName)
}

func (l *ComponentsLauncher) Launch() {
	for componentsName, cfg := range l.config {
		log.Infof("launch component: %s", componentsName)
		log.Infof("%s config: %v", componentsName, cfg)
		launchWrapper(cfg, componentsName)
	}

	for _, fn := range l.components {
		l.group.Run(fn)
	}

	if err := l.group.Wait(); err != nil {
		panic("launch components error: " + err.Error())
	}
}

func launchComponent[T any](cfg config.Value, initMethod func(cfg components.ConfigMap[*T]) (func() error, error)) {
	configs, err := cfg.Map()
	if err != nil {
		panic("get component config error: " + err.Error())
	}

	_, component := components.Load(configs, initMethod)
	if err := component.Start(); err != nil {
		panic("launch component error: " + err.Error())
	}
}

func (l *ComponentsLauncher) register(name string, fn func() error) {
	l.components[name] = fn
}
