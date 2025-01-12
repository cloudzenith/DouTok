package snowflakeutil

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetSnowflakeId(t *testing.T) {
	InitDefaultSnowflakeNode(1)
	require.NotEqual(t, 0, GetSnowflakeId())
}
