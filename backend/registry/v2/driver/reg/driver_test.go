package reg

import (
	"dkdk/model"
	"dkdk/test/test_config"
	"github.com/google/uuid"
	"github.com/opencontainers/go-digest"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func GetDriver(t *testing.T) (driver Driver) {
	driver, err := NewDriver(test_config.LocalTestServer, "", "", true)
	assert.NoError(t, err)
	return
}

func TestDriver_ListTags(t *testing.T) {
	driver := GetDriver(t)
	content := "hello dkdk!\n"
	repository := model.RepositoryBasic{
		Name:      "dkdk/hello-world",
		Reference: "v2",
	}
	err := driver.Upload(repository, "hello", strings.NewReader(content), digest.FromString(content), int64(len(content)))
	assert.NoError(t, err)
	tags, err := driver.ListTags(repository.Name)
	assert.NoError(t, err)
	log.Logger.Info(tags)
}

func TestDriver_Upload(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		driver := GetDriver(t)
		content := "hello dkdk!\n"
		repository := model.RepositoryBasic{
			Name:      "dkdk/" + uuid.New().String(),
			Reference: "v2",
		}
		err := driver.Upload(repository, "hello", strings.NewReader(content), digest.FromString(content), int64(len(content)))
		assert.NoError(t, err)
		tags, err := driver.ListTags(repository.Name)
		assert.NoError(t, err)
		assert.True(t, len(tags) > 0)
	})
}
