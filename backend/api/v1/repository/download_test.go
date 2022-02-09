package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/opencontainers/go-digest"
	"github.com/ssst0n3/dkdk/test/test_data"
	"github.com/ssst0n3/dkdk/test/test_db"
	"github.com/ssst0n3/lightweight_api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestDownloadFileCommon(t *testing.T) {
	assert.NoError(t, test_db.InitRepository())
	content := "Hello World!\n"
	dgs := digest.FromString(content)
	{
		// upload
		d, repository, err := GetRegDriverByRepositoryConfigId(test_data.RepositoryA.ID)
		assert.NoError(t, err)
		err = d.Upload(repository.RepositoryBasic, "hello.txt", strings.NewReader(content), dgs, int64(len(content)))
		assert.NoError(t, err)
	}
	{
		//download
		d, repository, err := GetPureDriverByRepositoryConfigId(test_data.RepositoryA.ID)
		assert.NoError(t, err)
		router.GET("/test/:digest", func(c *gin.Context) {
			err := DownloadFileCommon(c, d, repository)
			if err != nil {
				lightweight_api.HandleInternalServerError(c, err)
				return
			}
		})
		url := "/test/" + dgs.String()
		req, err := http.NewRequest(http.MethodGet, url, nil)
		w := lightweight_api.ObjectOperate(req, router)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, content, w.Body.String())
	}
}
