package oras

import (
	"fmt"
	"github.com/containerd/containerd/remotes"
	"github.com/ssst0n3/dkdk/registry/v2/service_provider/huaweicloud"
	"github.com/ssst0n3/dkdk/test/test_config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetDriver(t *testing.T, getResolveFunc func(resolver remotes.Resolver) remotes.Resolver) Driver {
	//customMediaType := "application/vnd.docker.container.image.v1+json"
	customMediaType := ""
	d, err := NewDriver(test_config.LocalTestServer, "", "", true, customMediaType, getResolveFunc)
	assert.NoError(t, err)
	return d
}

func TestDriver_List(t *testing.T) {
	d := GetDriver(t, nil)
	assert.NoError(t, d.Login())
	assert.NoError(t, d.List(fmt.Sprintf("%s/%s", test_config.LocalTestServer, "/dkdk/hello-world:v2")))
}

func TestDriver(t *testing.T) {
	t.Run("local registry", func(t *testing.T) {
		ref := "127.0.0.1:14005/library/hello:latest"
		d := GetDriver(t, nil)
		assert.NoError(t, d.Login())
		assert.NoError(t, d.Upload(ref, "testtest", []byte("hello dkdk")))
		assert.NoError(t, d.Download(ref, "testtest"))
	})
	t.Run("huaweicloud swr", func(t *testing.T) {
		ref := "swr.cn-south-1.myhuaweicloud.com/*"
		d := GetDriver(t, huaweicloud.NewResolver)
		assert.NoError(t, d.Login())
		assert.NoError(t, d.Upload(ref, "testtest", []byte("hello dkdk")))
		assert.NoError(t, d.Download(ref, "testtest"))
	})
}
