package main

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/pkg/db"
	"github.com/go-kratos/kratos/v2/log"
	"os"
)

func migrateDB(data *conf.Data, logger log.Logger) {
	client, err := db.NewDBClient(data)
	if err != nil {
		logger.Log(log.LevelError, "Failed to connect database.", err)
		os.Exit(1)
	}
	logger.Log(log.LevelInfo, "Database migration started.")
	// auto migrate your schema
	client.GetDB().AutoMigrate(&model.User{})

	logger.Log(log.LevelInfo, "Database migration completed.")
	// 迁移完成后，可以根据需要退出程序或者执行其他逻辑
	os.Exit(0)
}
