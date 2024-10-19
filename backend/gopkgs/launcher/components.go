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
	group  *gofer.Group
	config map[string]config.Value
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
		config: configMap,
	}
}

func launchWrapper(cfg config.Value, componentsName string) {
	switch componentsName {
	case "mysql":
		launchComponent(cfg, mysqlx.Init)
	case "redis":
		launchComponent(cfg, redisx.Init)
	case "minio":
		launchComponent(cfg, miniox.Init)
	case "etcd":
		launchComponent(cfg, etcdx.Init)
	case "consul":
		launchComponent(cfg, consulx.Init)
	case "rmqconsumer":
		launchComponent(cfg, rmqconsumerx.Init)
	case "rmqproducer":
		launchComponent(cfg, rmqproducerx.Init)
	default:
		panic("unknown components name: " + componentsName)
	}
}

func (l *ComponentsLauncher) Launch() {
	for componentsName, cfg := range l.config {
		log.Infof("launch component: %s", componentsName)
		log.Infof("%s config: %v", componentsName, cfg)
		launchWrapper(cfg, componentsName)
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
