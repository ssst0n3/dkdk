package v1

import (
	"dkdk/model"
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

func TestListResource(t *testing.T) {
	assert.NoError(t, test_db.InitDirectory())
	req, err := http.NewRequest(http.MethodGet, DirectoryResource.BaseRelativePath, nil)
	assert.NoError(t, err)
	test.Admin(t, req)
	w := lightweight_api.ObjectOperate(req, router)
	var directories []model.Directory
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &directories))
	assert.Equal(t, len(test_data.Directories), len(directories))
	assert.Equal(t, test_data.DirectoryA.ID, directories[0].ID)
	assert.Equal(t, test_data.DirectoryA.Filename, directories[0].Filename)
}

func TestListDirectoryUnderDir(t *testing.T) {
	assert.NoError(t, test_db.InitDirectory())
	assert.NoError(t, test_db.InitNode())
	t.Run("list root directories", func(t *testing.T) {
		{
			url := fmt.Sprintf("%s/list/%d", DirectoryResource.BaseRelativePath, 0)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			assert.NoError(t, err)
			test.Admin(t, req)
			w := lightweight_api.ObjectOperate(req, router)
			var directories []model.DirectoryNode
			assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &directories))
			assert.Equal(t, 1, len(directories))
			assert.Equal(t, test_data.DirectoryA.ID, directories[0].ID)
			assert.Equal(t, test_data.DirectoryA.Filename, directories[0].Filename)
			assert.Equal(t, node.Directory, directories[0].Type)
		}
	})
	t.Run("list sub directories", func(t *testing.T) {
		{
			url := fmt.Sprintf("%s/list/%d", DirectoryResource.BaseRelativePath, test_data.DirectoryA.ID)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			assert.NoError(t, err)
			test.Admin(t, req)
			w := lightweight_api.ObjectOperate(req, router)
			var directories []model.DirectoryNode
			assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &directories))
			assert.Equal(t, 1, len(directories))
			assert.Equal(t, test_data.DirectoryB.ID, directories[0].ID)
			assert.Equal(t, test_data.DirectoryB.Filename, directories[0].Filename)
			assert.Equal(t, node.Directory, directories[0].Type)
		}
	})
}

func TestGetDirectoryPath(t *testing.T) {
	assert.NoError(t, test_db.InitDirectory())
	assert.NoError(t, test_db.InitNode())
	url := fmt.Sprintf("%s/path/%d", DirectoryResource.BaseRelativePath, test_data.DirectoryB.ID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	assert.NoError(t, err)
	test.Admin(t, req)
	w := lightweight_api.ObjectOperate(req, router)
	var resp []model.DirectoryPathResponse
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, 2, len(resp))
	assert.Equal(t, test_data.DirectoryA.ID, resp[0].ID)
}
