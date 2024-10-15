package server

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/server/warmup"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/miniox"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/mysqlx"
)

func warmUp(params *Params) {
	warmup.CheckAndCreateFileRepoTables(mysqlx.GetDBClient(context.Background()), params.fileTableShardingConfig, params.dbShardingTablesConfig)
	warmup.CheckAndCreateMinioBucket(miniox.GetClient(context.Background()), params.dbShardingTablesConfig)
	warmup.InitMinioPublicDirectoryV2()
}
