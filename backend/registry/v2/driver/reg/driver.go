package reg

import (
	"bytes"
	"context"
	"dkdk/model"
	"github.com/docker/distribution"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/docker/docker/api/types"
	"github.com/genuinetools/reg/registry"
	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"io"
	"net/http"
)

type Driver struct {
	Registry *registry.Registry
}

func NewDriver(serverAddress, username, secret string, insecure bool) (driver Driver, err error) {
	cred := types.AuthConfig{
		Username:      username,
		ServerAddress: serverAddress,
	}
	if username == "" {
		cred.IdentityToken = secret
	} else {
		cred.Password = secret
	}
	r, err := registry.New(
		context.Background(),
		cred,
		registry.Opt{
			Domain:   "",
			Insecure: insecure,
			Debug:    true,
			SkipPing: true,
			NonSSL:   insecure,
			Timeout:  0,
			Headers:  nil,
		},
	)
	r.Client.Transport.(*registry.CustomTransport).Transport.(*registry.ErrorTransport).
		Transport.(*registry.BasicTransport).Transport.(*registry.TokenTransport).
		Transport.(*http.Transport).ForceAttemptHTTP2 = false
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	driver = Driver{
		Registry: r,
	}
	return
}

func (d Driver) Login() (err error) {
	return
}

func (d Driver) ListTags(repositoryName string) (tags []string, err error) {
	tags, err = d.Registry.Tags(context.Background(), repositoryName)
	awesome_error.CheckErr(err)
	return
}

func (d Driver) GetManifests(repository model.RepositoryBasic) (manifests schema2.Manifest, err error) {
	manifests, err = d.Registry.ManifestV2(context.Background(), repository.Name, repository.Reference)
	awesome_error.CheckErr(err)
	return
}

func (d Driver) Upload(repository model.RepositoryBasic, fileName string, content io.Reader, dsg digest.Digest, size int64) (err error) {
	manifestsOld, err := d.GetManifests(repository)
	if err != nil && manifestsOld.Versioned.SchemaVersion != 0 {
		awesome_error.CheckErr(err)
		return
	}
	err = d.Registry.UploadLayer(context.Background(), repository.Name, dsg, content)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	layer := distribution.Descriptor{
		MediaType: ocispec.MediaTypeImageLayer,
		Digest:    dsg,
		Size:      size,
		Annotations: map[string]string{
			ocispec.AnnotationTitle: fileName,
		},
	}
	configBytes := []byte("{}")
	err = d.Registry.UploadLayer(context.Background(), repository.Name, digest.FromBytes(configBytes), bytes.NewReader(configBytes))
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	manifest, err := schema2.FromStruct(schema2.Manifest{
		Versioned: schema2.SchemaVersion,
		Config: distribution.Descriptor{
			MediaType: schema2.MediaTypeImageConfig,
			Size:      int64(len(configBytes)),
			Digest:    digest.FromBytes(configBytes),
		},
		Layers: append(manifestsOld.Layers, layer),
	})
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	err = d.Registry.PutManifest(context.Background(), repository.Name, repository.Reference, manifest)
	awesome_error.CheckErr(err)
	return
}

func (d Driver) GetFileDigestSizeByFilename(repository model.RepositoryBasic, filename string) (dsg digest.Digest, size int64, err error) {
	manifest, err := d.GetManifests(repository)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	for _, layer := range manifest.Layers {
		if layer.Annotations[ocispec.AnnotationTitle] == filename {
			dsg = layer.Digest
			size = layer.Size
			break
		}
	}
	return
}

func (d Driver) GetSizeByDigest(repository model.RepositoryBasic, digest string) (size int64, err error) {
	manifest, err := d.GetManifests(repository)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	for _, layer := range manifest.Layers {
		if layer.Digest.String() == digest {
			size = layer.Size
			break
		}
	}
	return
}

func (d Driver) Download(repository model.RepositoryBasic, identifier model.FileIdentifier) (content io.Reader, size int64, err error) {
	if identifier.Digest == "" {
		identifier.Digest, size, err = d.GetFileDigestSizeByFilename(repository, identifier.Filename)
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
	} else {
		size, err = d.GetSizeByDigest(repository, identifier.Digest.String())
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}

	}
	content, err = d.Registry.DownloadLayer(context.Background(), repository.Name, identifier.Digest)
	awesome_error.CheckErr(err)
	return
}

func (d Driver) Delete(repository model.RepositoryBasic, identifier model.FileIdentifier) (err error) {
	// TODO
	return
}
