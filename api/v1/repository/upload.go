package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/opencontainers/go-digest"
	"github.com/ssst0n3/dkdk/model"
	"github.com/ssst0n3/dkdk/registry/v2/driver"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/response"
)

func uploadFileCommon(c *gin.Context, d driver.Driver, repository model.Repository) (err error) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}

	file2, err := fileHeader.Open()
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	dgs, err := digest.FromReader(file2)
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	// TODO: check digest is already exists
	err = d.Upload(
		model.RepositoryBasic{Name: repository.Name, Reference: repository.Reference},
		fileHeader.Filename,
		file,
		dgs,
		fileHeader.Size,
	)
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	response.Success200(c, "upload success")
	return
}

func UploadFileByRepositoryId(c *gin.Context) {
	repositoryId, err := Resource.MustResourceExistsByIdAutoParseParam(c)
	if err != nil {
		return
	}
	d, repository, err := GetRegDriverByRepositoryConfigId(uint(repositoryId))
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	_ = uploadFileCommon(c, d, repository)
}

func UploadFileFromRepositoryConfig(c *gin.Context) {
	d, repository, err := GetDriverAutoParseRepositoryConfig(c)
	if err != nil {
		return
	}
	_ = uploadFileCommon(c, d, model.Repository{
		Registry:        repository.Registry,
		Cred:            repository.Cred,
		RepositoryBasic: repository.RepositoryBasic,
	})
}
