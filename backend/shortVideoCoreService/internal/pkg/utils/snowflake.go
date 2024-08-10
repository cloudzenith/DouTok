package utils

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
)

type SnowflakeNode struct {
	node *snowflake.Node
}

func (s *SnowflakeNode) GetSnowflakeId() int64 {
	return s.node.Generate().Int64()
}

func NewSnowflakeNode(c *conf.Common) (*SnowflakeNode, error) {
	defaultNode, err := snowflake.NewNode(c.Node)
	if err != nil {
		return nil, err
	}
	return &SnowflakeNode{node: defaultNode}, nil
}
