package test

import (
	"dkdk/model"
	"dkdk/registry/v2/driver"
	"dkdk/registry/v2/driver/pure"
	"dkdk/registry/v2/driver/reg"
	"dkdk/test/test_config"
	"github.com/opencontainers/go-digest"
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func DriverTestSuite(t *testing.T, driver driver.Driver) {
	repository := model.RepositoryBasic{
		Name:      "dkdk/hello-world",
		Reference: "v2",
	}
	content := "hello dkdk!\n"
	{
		// UPLOAD
		err := driver.Upload(repository, "hello", strings.NewReader(content), digest.FromString(content), int64(len(content)))
		assert.NoError(t, err)
	}
	{
		// DOWNLOAD
		download, _, err := driver.Download(repository, model.FileIdentifier{Filename: "hello"})
		assert.NoError(t, err)
		bytes, err := ioutil.ReadAll(download)
		assert.NoError(t, err)
		assert.Equal(t, content, string(bytes))
	}
	{
		// LIST TAG
		tags, err := driver.ListTags(repository.Name)
		assert.NoError(t, err)
		log.Logger.Info(tags)
	}
	{
		// GET Manifests
		tags, err := driver.GetManifests(repository)
		assert.NoError(t, err)
		log.Logger.Info(tags)
	}
	{
		// DELETE
		assert.NoError(t, driver.Delete(repository, model.FileIdentifier{Filename: "hello"}))
	}
}

func TestReg(t *testing.T) {
	t.Run("local", func(t *testing.T) {
		d, err := reg.NewDriver(test_config.LocalTestServer, "", "", true)
		assert.NoError(t, err)
		DriverTestSuite(t, d)
	})
}

func TestPure(t *testing.T) {
	t.Run("list tags", func(t *testing.T) {
		log.Logger.Level = logrus.DebugLevel
		d := pure.NewDriver(test_config.LocalTestServer, "", "", true)
		tags, err := d.ListTags("dkdk/hello-world")
		assert.NoError(t, err)
		log.Logger.Info(tags)
	})
}
