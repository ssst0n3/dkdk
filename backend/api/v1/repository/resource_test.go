package repository

import (
	"bytes"
	"dkdk/database"
	"dkdk/model"
	"dkdk/test"
	"dkdk/test/test_data"
	"dkdk/test/test_db"
	"encoding/json"
	"fmt"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/lightweight_api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestListRepository(t *testing.T) {
	assert.NoError(t, test_db.InitRepository())
	router.GET(Resource.BaseRelativePath, ListRepository)
	req, err := http.NewRequest(http.MethodGet, Resource.BaseRelativePath, nil)
	assert.NoError(t, err)
	test.Admin(t, req)
	w := lightweight_api.ObjectOperate(req, router)
	log.Logger.Infof("body: %+v", w.Body.String())
}

func TestUpdateRepository(t *testing.T) {
	assert.NoError(t, test_db.InitRepository())
	router.PUT("/common/:id", UpdateRepository)
	r := model.RepositoryConfigUpdateCommonBody{
		Username: "test",
	}
	body, err := json.Marshal(r)
	assert.NoError(t, err)
	req, err := http.NewRequest(
		http.MethodPut, fmt.Sprintf("/common/%d", test_data.RepositoryA.ID), bytes.NewReader(body),
	)
	assert.NoError(t, err)
	test.Admin(t, req)

	assert.NotEqual(t, r.Username, test_data.RepositoryA.Username)
	w := lightweight_api.ObjectOperate(req, router)
	assert.Equal(t, http.StatusOK, w.Code)
	var query model.Repository
	assert.NoError(t, database.DB.First(&query, test_data.RepositoryA.ID).Error)
	assert.Equal(t, r.Username, query.Username)
}
