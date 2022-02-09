package driver

import (
	"dkdk/model"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/opencontainers/go-digest"
	"io"
)

type Driver interface {
	Login() (err error)
	ListTags(repositoryName string) (tags []string, err error)
	GetManifests(repository model.RepositoryBasic) (manifests schema2.Manifest, err error)
	Upload(repository model.RepositoryBasic, filename string, content io.Reader, dsg digest.Digest, size int64) (err error)
	Download(repository model.RepositoryBasic, identifier model.FileIdentifier) (content io.Reader, size int64, err error)
	Delete(repository model.RepositoryBasic, identifier model.FileIdentifier) (err error)
}
