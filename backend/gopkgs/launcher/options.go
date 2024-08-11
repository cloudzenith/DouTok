package launcher

import "github.com/go-kratos/kratos/v2/config"

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
