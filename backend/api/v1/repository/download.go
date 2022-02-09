package repository

import (
	"dkdk/database"
	"dkdk/model"
	"dkdk/registry/v2/driver"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opencontainers/go-digest"
	"github.com/ssst0n3/lightweight_api"
	"net/http"
)

func DownloadFileCommon(c *gin.Context, d driver.Driver, repository model.Repository) (err error) {
	dsg := digest.Digest(c.Param("digest"))
	content, size, err := d.Download(
		model.RepositoryBasic{Name: repository.Name, Reference: repository.Reference},
		model.FileIdentifier{Digest: dsg},
	)
	if err != nil {
		return
	}
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.png"`,
	}
	c.DataFromReader(http.StatusOK, size, "application/octet-stream", content, extraHeaders)
	err = database.UpdateFileDownloadCount(dsg)
	if err != nil {
		return
	}
	return
}

func DownloadFileByRepositoryId(c *gin.Context) {
	userId, err := lightweight_api.GetUserId(c)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	repositoryId, err := Resource.MustResourceExistsByIdAutoParseParam(c)
	if err != nil {
		return
	}
	if has, err := database.CheckUserHasRepository(userId, uint(repositoryId)); err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	} else if !has {
		err := errors.New(fmt.Sprintf("You do not have repository %d", repositoryId))
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	d, repository, err := GetPureDriverByRepositoryConfigId(uint(repositoryId))
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	err = DownloadFileCommon(c, d, repository)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
}

func DownloadFileFromRepositoryConfig(c *gin.Context) {
	d, repository, err := GetDriverAutoParseRepositoryConfig(c)
	if err != nil {
		return
	}
	err = DownloadFileCommon(c, d, model.Repository{
		Registry:        repository.Registry,
		Cred:            repository.Cred,
		RepositoryBasic: repository.RepositoryBasic,
	})
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
}
