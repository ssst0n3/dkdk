package offline_download

import (
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestArchive(t *testing.T) {
	path := "/tmp/test.txt"
	assert.NoError(t, ioutil.WriteFile(path, []byte("data"), 0755))
	outputPath, err := Archive(path, "test.txt", "123")
	assert.NoError(t, err)
	log.Logger.Info(outputPath)
}
