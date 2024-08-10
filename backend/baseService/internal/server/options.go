package server

import (
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/service/fileservice"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type Params struct {
	addr                    string
	redisDsn                string
	redisPassword           string
	db                      *gorm.DB
	minioCore               *minio.Core
	fileTableShardingConfig fileservice.FileTableShardingConfig
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

func WithDB(db *gorm.DB) Option {
	return func(p *Params) {
		p.db = db
	}
}

func WithMinioCore(core *minio.Core) Option {
	return func(p *Params) {
		p.minioCore = core
	}
}

func WithFileTableShardingConfig(config fileservice.FileTableShardingConfig) Option {
	return func(p *Params) {
		p.fileTableShardingConfig = config
	}
}
