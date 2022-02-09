package util

import (
	"github.com/docker/distribution/manifest/schema2"
	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/dkdk/model"
	"github.com/ssst0n3/dkdk/registry/v2/driver"
)

func GetSizeByDigestFromManifest(digest string, manifest schema2.Manifest) (size int64, exists bool) {
	for _, layer := range manifest.Layers {
		if layer.Digest.String() == digest {
			size = layer.Size
			exists = true
			break
		}
	}
	return
}

func GetFileDigestSizeByFilenameFromManifest(filename string, manifest schema2.Manifest) (dsg digest.Digest, size int64, exists bool) {
	for _, layer := range manifest.Layers {
		if layer.Annotations[ocispec.AnnotationTitle] == filename {
			dsg = layer.Digest
			size = layer.Size
			exists = true
			break
		}
	}
	return
}

func ListFile(d driver.Driver, repository model.RepositoryBasic) (fileItems []model.FileItem, err error) {
	manifests, err := d.GetManifests(repository)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	for _, layer := range manifests.Layers {
		filename := layer.Annotations[ocispec.AnnotationTitle]
		fileItems = append(fileItems, model.FileItem{
			FileIdentifier: model.FileIdentifier{
				Filename: filename,
				Digest:   layer.Digest,
			},
			Size: layer.Size,
		})
	}
	return
}
