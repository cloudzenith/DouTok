package db

import (
	"fmt"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBClient struct {
	db *gorm.DB
}

func NewDBClient(c *conf.Data, opts ...gorm.Option) (*DBClient, error) {
	dsn := c.Database.Source
	if c.Database.Driver == "mysql" {
		db, err := gorm.Open(mysql.Open(dsn), opts...)
		if err != nil {
			return nil, err
		}
		return &DBClient{db: db}, nil
	}
	return nil, nil
}

func (c *DBClient) GetDB() *gorm.DB {
	return c.db
}

func (c *DBClient) StartTransaction() (*gorm.DB, *TransactionMaker, error) {
	if c.db == nil {
		return nil, nil, fmt.Errorf("database not provided")
	}
	tx := c.db.Begin()
	return tx, &TransactionMaker{tx: tx}, nil
}
