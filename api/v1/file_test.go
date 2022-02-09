package v1

import (
	"encoding/json"
	"github.com/ssst0n3/dkdk/model"
	"github.com/ssst0n3/dkdk/model/node"
	"github.com/ssst0n3/dkdk/test"
	"github.com/ssst0n3/dkdk/test/test_data"
	"github.com/ssst0n3/dkdk/test/test_db"
	"github.com/ssst0n3/lightweight_api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestListFileUnderDir(t *testing.T) {
	assert.NoError(t, test_db.InitFile())
	t.Run("empty", func(t *testing.T) {
		assert.NoError(t, test_db.MakeNodeEmpty())
		req, err := http.NewRequest(http.MethodGet, FileResource.BaseRelativePath+"/0", nil)
		assert.NoError(t, err)
		test.Admin(t, req)
		w := lightweight_api.ObjectOperate(req, router)
		assert.Equal(t, "[]", w.Body.String())
	})
	assert.NoError(t, test_db.InitNode())
	t.Run("list root files", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, FileResource.BaseRelativePath+"/0", nil)
		assert.NoError(t, err)
		test.Admin(t, req)
		w := lightweight_api.ObjectOperate(req, router)
		var files []model.FileNode
		assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &files))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, test_data.FileA.ID, files[0].ID)
		assert.Equal(t, test_data.FileA.Filename, files[0].Filename)
		assert.Equal(t, node.File, files[0].Type)
	})
	t.Run("list sub files", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, FileResource.BaseRelativePath+"/"+strconv.Itoa(int(test_data.FileA.ID)), nil)
		assert.NoError(t, err)
		test.Admin(t, req)
		w := lightweight_api.ObjectOperate(req, router)
		var files []model.FileNode
		assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &files))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, test_data.FileB.ID, files[0].ID)
		assert.Equal(t, test_data.FileB.Filename, files[0].Filename)
		assert.Equal(t, node.File, files[0].Type)
	})
}

func TestUploadRepositoryFileToDirectory(t *testing.T) {
	// TODO
}
