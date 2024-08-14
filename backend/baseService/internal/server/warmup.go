package server

import "github.com/cloudzenith/DouTok/backend/baseService/internal/server/warmup"

func warmUp(params *Params) {
	warmup.CheckAndCreateFileRepoTables(params.db, params.fileTableShardingConfig, params.dbShardingTablesConfig)
	warmup.CheckAndCreateMinioBucket(params.minioCore, params.dbShardingTablesConfig)
}
