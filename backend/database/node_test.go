package database

import (
	"github.com/ssst0n3/dkdk/model/node"
	"github.com/ssst0n3/dkdk/test/test_data"
	"github.com/ssst0n3/dkdk/test/test_db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBatchDeleteNode(t *testing.T) {
	assert.NoError(t, test_db.InitNode())
	var count int64
	DB.Model(&node.Node{}).Count(&count)
	assert.Equal(t, len(test_data.Nodes), int(count))
	var ids []uint
	for _, n := range test_data.Nodes {
		ids = append(ids, n.ID)
	}
	assert.NoError(t, BatchDeleteNode(ids))
	DB.Model(&node.Node{}).Count(&count)
	assert.Equal(t, 0, int(count))
}
