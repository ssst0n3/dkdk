package repository

import (
	"dkdk/database"
	"dkdk/model"
	"dkdk/registry/v2/driver/reg"
	"github.com/gin-gonic/gin"
	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/ssst0n3/lightweight_api"
	"net/http"
)

func ListFileFromAllRepository(c *gin.Context) {
	userId, err := lightweight_api.GetUserId(c)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	repositoryList, err := database.ListRepositoryByUserId(userId)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	for _, repository := range repositoryList {
		driver, err := reg.NewDriver(repository.ServiceAddress, repository.Username, repository.Secret, repository.Insecure)
		if err != nil {
			lightweight_api.HandleInternalServerError(c, err)
			return
		}
		// TODO
		driver.GetManifests(repository.RepositoryBasic)
	}

}

// ListFileByRepositoryId godoc
// @Summary list layer of a repository
// @Description return layers
// @Tags Repository
// @ID list-file-by-repository-id
// @Accept json
// @Produce json
// @Param id path int true "Repository ID"
// @Success 200 {array} model.FileItem
// @Router /api/v1/repository/list/{id} [get]
func ListFileByRepositoryId(c *gin.Context) {
	repository, err := GetRepositoryFromParamId(c)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	driver, err := reg.NewDriver(repository.ServiceAddress, repository.Username, repository.Secret, repository.Insecure)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	manifests, err := driver.GetManifests(repository.RepositoryBasic)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	var digests []digest.Digest
	var fileItems []model.ResponseListFileUnderRepository
	for _, layer := range manifests.Layers {
		filename := layer.Annotations[ocispec.AnnotationTitle]
		digests = append(digests, layer.Digest)
		fileItems = append(fileItems, model.ResponseListFileUnderRepository{
			FileItem: model.FileItem{
				FileIdentifier: model.FileIdentifier{
					Filename: filename,
					Digest:   layer.Digest,
				},
				Size: layer.Size,
			},
		})
	}
	exists, err := database.DigestExists(digests)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	for i, fileItem := range fileItems {
		if _, ok := exists[fileItem.Digest]; ok {
			fileItem.InNetDisk = true
			//fileItem.FilenameInDkdk = filename
			fileItems[i] = fileItem
		}
	}
	c.JSON(http.StatusOK, fileItems)
}

// ListFileFromRepositoryConfig godoc
// @Summary list layer of a repository
// @Description return layers
// @Tags Repository
// @ID list-file-from-repository-config
// @Param repository body model.Repository true "List File By Repository"
// @Accept json
// @Produce json
// @Success 200 {array} model.FileItem
// @Router /api/v1/repository/list [post]
func ListFileFromRepositoryConfig(c *gin.Context) {
	driver, repository, err := GetDriverAutoParseRepositoryConfig(c)
	if err != nil {
		return
	}
	manifests, err := driver.GetManifests(repository.RepositoryBasic)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	var fileItems []model.FileItem
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
	c.JSON(http.StatusOK, fileItems)
}
