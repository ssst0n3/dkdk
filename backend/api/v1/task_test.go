package v1

import (
	"bytes"
	"dkdk/test"
	"dkdk/test/test_data"
	"encoding/json"
	"github.com/ssst0n3/lightweight_api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestBatchCreateTask(t *testing.T) {
	task := test_data.Task1
	body, err := json.Marshal(task)
	assert.NoError(t, err)
	req, err := http.NewRequest(http.MethodPost, TaskResource.BaseRelativePath, bytes.NewReader(body))
	assert.NoError(t, err)
	test.Admin(t, req)
	w := lightweight_api.ObjectOperate(req, router)
	assert.Equal(t, http.StatusOK, w.Code)
}

