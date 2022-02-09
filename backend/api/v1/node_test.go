package v1

import (
	"bytes"
	"dkdk/database"
	"dkdk/model/node"
	"dkdk/test"
	"dkdk/test/test_data"
	"dkdk/test/test_db"
	"encoding/json"
	"fmt"
	"github.com/ssst0n3/lightweight_api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestBatchDeleteNode(t *testing.T) {
	assert.NoError(t, test_db.InitNode())
	var count int64
	database.DB.Model(&node.Node{}).Count(&count)
	assert.Equal(t, len(test_data.Nodes), int(count))
	var ids []uint
	for _, n := range test_data.Nodes {
		ids = append(ids, n.ID)
	}
	body, err := json.Marshal(ids)
	assert.NoError(t, err)
	req, err := http.NewRequest(http.MethodDelete, NodeResource.BaseRelativePath, bytes.NewReader(body))
	assert.NoError(t, err)
	test.Admin(t, req)
	w := lightweight_api.ObjectOperate(req, router)
	assert.Equal(t, http.StatusOK, w.Code)
	database.DB.Model(&node.Node{}).Count(&count)
	assert.Equal(t, 0, int(count))
}

func TestMove(t *testing.T) {
	assert.NoError(t, test_db.InitNode())
	body, err := json.Marshal(node.MoveBody{Parent: test_data.NodeA.Parent})
	assert.NoError(t, err)
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%d", NodeResource.BaseRelativePath, test_data.NodeC.ID), bytes.NewReader(body))
	assert.NoError(t, err)
	test.Admin(t, req)
	w := lightweight_api.ObjectOperate(req, router)
	assert.Equal(t, http.StatusOK, w.Code)
	n := node.Node{}
	assert.NoError(t, database.DB.Model(&node.Node{}).First(&n, test_data.NodeC.ID).Error)
	assert.Equal(t, test_data.NodeA.Parent, n.Parent)
}
