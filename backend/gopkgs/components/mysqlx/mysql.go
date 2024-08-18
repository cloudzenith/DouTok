package mysqlx

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

var (
	globalClientMap = sync.Map{}
	globalConfigMap = make(components.ConfigMap[*Config])
)

func GetConfig() components.ConfigMap[*Config] {
	return globalConfigMap
}

func Init(cm components.ConfigMap[*Config]) error {
	globalConfigMap = cm

	for k, v := range cm {
		db, err := Connect(v)

		if err != nil {
			return err
		}

		globalClientMap.Store(k, db)
	}

	return nil
}

func Connect(c *Config) (*gorm.DB, error) {
	c.SetDefault()
	originDB, err := sql.Open("mysql", c.ToDSN())
	if err != nil {
		return nil, err
	}

	originDB.SetMaxIdleConns(c.MaxIdle)
	originDB.SetMaxOpenConns(c.MaxOpen)

	if c.ConnMaxLifeTime > 0 {
		originDB.SetConnMaxLifetime(time.Duration(c.ConnMaxLifeTime) * time.Second)
	} else {
		originDB.SetConnMaxLifetime(0)
	}

	if c.ConnMaxIdleTime > 0 {
		originDB.SetConnMaxIdleTime(time.Duration(c.ConnMaxIdleTime) * time.Second)
	} else {
		originDB.SetConnMaxIdleTime(0)
	}

	var connPoll gorm.ConnPool = originDB

	dialector := mysql.New(mysql.Config{
		Conn: connPoll,
	})

	return gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
}

func getKey(keys ...string) string {
	if len(keys) == 0 {
		return "default"
	}

	return keys[0]
}

func GetDBClient(ctx context.Context, keys ...string) *gorm.DB {
	key := getKey(keys...)

	value, ok := globalClientMap.Load(key)
	if !ok {
		panic(fmt.Sprintf("%s not init", key))
	}

	return value.(*gorm.DB)
}
