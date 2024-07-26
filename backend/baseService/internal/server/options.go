package server

type Params struct {
	addr          string
	redisDsn      string
	redisPassword string
}

type Option func(*Params)

func WithAddr(addr string) Option {
	return func(p *Params) {
		p.addr = addr
	}
}

func WithRedisDsn(redisDsn string) Option {
	return func(p *Params) {
		p.redisDsn = redisDsn
	}
}

func WithRedisPassword(redisPassword string) Option {
	return func(p *Params) {
		p.redisPassword = redisPassword
	}
}
