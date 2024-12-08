package launcher

import (
	"context"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/consulx"
	"github.com/cloudzenith/DouTok/backend/gopkgs/internal/defaultlogger"
	"github.com/cloudzenith/DouTok/backend/gopkgs/internal/shutdown"
	"github.com/cloudzenith/DouTok/backend/gopkgs/snowflakeutil"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"os"
	"sync"
	"time"
)

type Launcher struct {
	app                     *kratos.App
	notNeedServiceDiscovery bool

	configOptions  []config.Option
	configWatchMap map[string]config.Observer
	config         config.Config
	configValue    interface{}

	logger        log.Logger
	grpcServer    func(configValue interface{}) *grpc.Server
	ginServer     func(configValue interface{}) *http.Server
	kratosOptions []kratos.Option

	componentsLauncher *ComponentsLauncher

	beforeConfigInitHandlers  []func()
	afterConfigInitHandlers   []func()
	beforeServerStartHandlers []func()
	afterServerStartHandlers  []func()
	shutdownHandlers          []func()
}

func New(options ...Option) *Launcher {
	launcher := &Launcher{
		configWatchMap: make(map[string]config.Observer),
	}
	for _, option := range options {
		option(launcher)
	}

	return launcher
}

func (l *Launcher) Run() {
	dir, _ := os.Getwd()
	log.Context(context.Background()).Info("current work directory: %s", dir)

	l.runInitConfig()

	l.runHandlers(l.beforeServerStartHandlers, "start to run handlers before server start")
	l.componentsLauncher.Launch()
	l.newKratosApp()
	<-l.run()
	l.runHandlers(l.afterServerStartHandlers, "start to run handlers after server start")

	<-shutdown.FiredCh()
	shutdown.Wait(10 * time.Second)
	l.runHandlers(l.shutdownHandlers, "start to run shutdown handlers")
}

func (l *Launcher) runInitConfig() {
	l.runHandlers(l.beforeConfigInitHandlers, "start to run handlers before config init")

	cfg := config.New(
		l.configOptions...,
	)
	if err := cfg.Load(); err != nil {
		panic(fmt.Errorf("failed to load config: %v", err))
	}

	l.config = cfg
	for key, observer := range l.configWatchMap {
		if err := cfg.Watch(key, observer); err != nil {
			panic(fmt.Errorf("failed to watch config: %v", err))
		}
	}

	if err := cfg.Scan(l.configValue); err != nil {
		panic(fmt.Errorf("failed to scan config value: %v", err))
	}

	l.componentsLauncher = NewComponentsLauncher(cfg)
	l.runHandlers(l.afterConfigInitHandlers, "start to run handlers after config init")
}

func (l *Launcher) initTracer(cfg *App) {
	if cfg.TraceEndpoint != "" {
		if err := initTracer(cfg.Name, cfg.TraceEndpoint); err != nil {
			panic(err)
		}
	}
}

func (l *Launcher) runHandlers(handlers []func(), info string) {
	if len(handlers) > 0 {
		log.Context(context.Background()).Info(info)
	}

	for _, handler := range handlers {
		handler()
	}
}

func (l *Launcher) newKratosApp() {
	options := make([]kratos.Option, 0)

	if l.logger != nil {
		options = append(options, kratos.Logger(l.logger))
	} else {
		options = append(options, kratos.Logger(defaultlogger.GetLogger()))
	}

	if l.grpcServer != nil {
		options = append(options, kratos.Server(l.grpcServer(l.configValue)))
	}

	if l.ginServer != nil {
		options = append(options, kratos.Server(l.ginServer(l.configValue)))
	}

	if len(l.kratosOptions) > 0 {
		options = append(options, l.kratosOptions...)
	}

	if !l.notNeedServiceDiscovery {
		consulClient := consulx.GetClient(context.Background())
		consulReg := consul.New(consulClient)
		options = append(options, kratos.Registrar(consulReg))
	}

	value := l.config.Value("app")
	appConfig := &App{}
	if err := value.Scan(appConfig); err != nil {
		panic(fmt.Errorf("failed to scan app config: %v", err))
	}
	options = append(
		options, kratos.Name(appConfig.Name), kratos.Version(appConfig.Version),
	)

	l.app = kratos.New(options...)

	l.initTracer(appConfig)
}

// nolint
func (l *Launcher) initSnowflake(appConfig *App) {
	if appConfig.Node == 0 {
		panic("snowflake node should be set")
	}

	snowflakeutil.InitDefaultSnowflakeNode(appConfig.Node)
}

func (l *Launcher) runKratosApp() <-chan struct{} {
	if l.app == nil {
		panic("app not initialized")
	}

	readyChan := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wg.Done()

		if err := l.app.Run(); err != nil {
			log.Context(context.Background()).Fatal("failed to run app")
			panic(err)
		}
	}()

	go func() {
		wg.Wait()
		close(readyChan)
	}()

	return readyChan
}

func (l *Launcher) run() chan struct{} {
	ch := l.runKratosApp()
	appReadyCh := make(chan struct{})

	go func() {
		<-ch
		close(appReadyCh)
	}()

	return appReadyCh
}

func (l *Launcher) ScanConfig(v interface{}) error {
	return l.config.Scan(v)
}
