package pure

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/opencontainers/go-digest"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/dkdk/model"
	"github.com/ssst0n3/dkdk/registry/v2/driver/util"
	"github.com/ssst0n3/registry_v2_client/registry"
	"io"
)

type Driver struct {
	Registry registry.Registry
}

func NewDriver(serviceAddress, username, password string, insecure bool) Driver {
	return Driver{Registry: registry.NewRegistry(serviceAddress, username, password, insecure)}
}

func (d Driver) Login() (err error) {
	return
}

func (d Driver) ListTags(repositoryName string) (tags []string, err error) {
	tags, err = d.Registry.GetTags(repositoryName)
	return
}

func (d Driver) GetManifests(repository model.RepositoryBasic) (manifests schema2.Manifest, err error) {
	manifests, err = d.Registry.GetManifest(repository.Name, repository.Reference)
	return
}

func (d Driver) Upload(repository model.RepositoryBasic, filename string, content io.Reader, dsg digest.Digest, size int64) (err error) {
	return
}

func (d Driver) Download(repository model.RepositoryBasic, identifier model.FileIdentifier) (content io.Reader, size int64, err error) {
	manifest, err := d.GetManifests(repository)
	if err != nil {
		return
	}
	if identifier.Digest == "" {
		dgs, _, exists := util.GetFileDigestSizeByFilenameFromManifest(identifier.Filename, manifest)
		if !exists {
			err = errors.New(fmt.Sprintf("file %s not exists", identifier.Filename))
			awesome_error.CheckErr(err)
			return
		}
		identifier.Digest = dgs
	} else {
		_, exists := util.GetSizeByDigestFromManifest(identifier.Digest.String(), manifest)
		if !exists {
			err = errors.New(fmt.Sprintf("file %s not exists", identifier.Digest))
			awesome_error.CheckErr(err)
			return
		}
	}
	data, err := d.Registry.FetchBlob(repository.Name, identifier.Digest.String(), false)
	size = int64(len(data))
	if err != nil {
		return
	}
	content = bytes.NewBuffer(data)
	return
}

func (d Driver) Delete(repository model.RepositoryBasic, identifier model.FileIdentifier) (err error) {
	return
}
