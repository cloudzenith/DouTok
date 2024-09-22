package warmup

import (
	"context"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/innerservice/filerepohelper"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

func CheckAndCreateFileRepoTables(
	db *gorm.DB,
	fileTableShardingConfig filerepohelper.FileTableShardingConfig,
	shardingConfigs map[string]conf.DomainShardingConfig,
) {
	for tableName, domainConfig := range shardingConfigs {
		iterTable(db, tableName, domainConfig, fileTableShardingConfig)
	}
}

func iterTable(db *gorm.DB, tableName string, domainConfig conf.DomainShardingConfig, fileTableShardingConfig filerepohelper.FileTableShardingConfig) {
	for bizName, bizConfig := range domainConfig.BizShardingFieldConfig {
		iterBiz(db, tableName, domainConfig.DomainName, bizName, bizConfig.Fields, fileTableShardingConfig)
	}
}

func iterBiz(db *gorm.DB, tableName, domainName, bizName string, fieldNameList []string, fileTableShardingConfig filerepohelper.FileTableShardingConfig) {
	for _, fieldName := range fieldNameList {
		shardingNum := fileTableShardingConfig.GetShardingNumber(tableName, domainName, bizName)
		for i := 0; i < int(shardingNum); i++ {
			shardingTableName := fmt.Sprintf("%s_%s_%s_%s_%d", tableName, domainName, bizName, fieldName, i)
			checkAndCreateTable(db, shardingTableName, tableName)
		}
	}

}

func checkAndCreateTable(db *gorm.DB, tableName, templateTableName string) {
	log.Context(context.Background()).Infof("check and create table %s, source table is %s", tableName, templateTableName)
	isExists := db.Migrator().HasTable(tableName)
	if isExists {
		log.Context(context.Background()).Infof("table %s already exists", tableName)
		return
	}

	if err := db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` LIKE `%s`", tableName, templateTableName)).Error; err != nil {
		log.Context(context.Background()).Errorf("create table %s failed: %v", tableName, err)
		panic(err)
	}
}
