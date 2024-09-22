package server

import (
	"github.com/cloudzenith/DouTok/backend/baseService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/innerservice/filerepohelper"
)

type Params struct {
	addr                    string
	fileTableShardingConfig filerepohelper.FileTableShardingConfig
	dbShardingTablesConfig  map[string]conf.DomainShardingConfig
}

type Option func(*Params)

func WithAddr(addr string) Option {
	return func(p *Params) {
		p.addr = addr
	}
}

func WithFileTableShardingConfig(config filerepohelper.FileTableShardingConfig) Option {
	return func(p *Params) {
		p.fileTableShardingConfig = config
	}
}

func WithDBShardingTablesConfig(config map[string]conf.DomainShardingConfig) Option {
	return func(p *Params) {
		p.dbShardingTablesConfig = config
	}
}
