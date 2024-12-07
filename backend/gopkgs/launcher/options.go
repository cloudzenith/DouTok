package launcher

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type Option func(*Launcher)

func WithBeforeConfigInitHandler(handler func()) Option {
	return func(l *Launcher) {
		l.beforeConfigInitHandlers = append(l.beforeConfigInitHandlers, handler)
	}
}

func WithAfterConfigInitHandler(handler func()) Option {
	return func(l *Launcher) {
		l.afterConfigInitHandlers = append(l.afterConfigInitHandlers, handler)
	}
}

func WithBeforeServerStartHandler(handler func()) Option {
	return func(l *Launcher) {
		l.beforeServerStartHandlers = append(l.beforeServerStartHandlers, handler)
	}
}

func WithAfterServerStartHandler(handler func()) Option {
	return func(l *Launcher) {
		l.afterServerStartHandlers = append(l.afterServerStartHandlers, handler)
	}
}

func WithShutdownHandler(handler func()) Option {
	return func(l *Launcher) {
		l.shutdownHandlers = append(l.shutdownHandlers, handler)
	}
}

func WithConfigOptions(options ...config.Option) Option {
	return func(l *Launcher) {
		l.configOptions = append(l.configOptions, options...)
	}
}

func WithConfigWatcher(key string, observer config.Observer) Option {
	return func(l *Launcher) {
		l.configWatchMap[key] = observer
	}
}

func WithHttpServer(s func(configValue interface{}) *http.Server) Option {
	return func(l *Launcher) {
		l.ginServer = s
	}
}

func WithGrpcServer(s func(configValue interface{}) *grpc.Server) Option {
	return func(l *Launcher) {
		l.grpcServer = s
	}
}

func WithConfigValue(value interface{}) Option {
	return func(l *Launcher) {
		l.configValue = value
	}
}

func WithKratosOptions(options ...kratos.Option) Option {
	return func(l *Launcher) {
		l.kratosOptions = append(l.kratosOptions, options...)
	}
}

func WithLogger(logger log.Logger) Option {
	return func(l *Launcher) {
		l.logger = logger
	}
}

func WithoutServiceDiscovery() Option {
	return func(l *Launcher) {
		l.notNeedServiceDiscovery = true
	}
}
