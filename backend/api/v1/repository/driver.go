package repository

import (
	"dkdk/database"
	"dkdk/model"
	"dkdk/registry/v2/driver"
	"dkdk/registry/v2/driver/pure"
	"dkdk/registry/v2/driver/reg"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ssst0n3/lightweight_api"
	"net/http"
)

func GetPureDriverByRepositoryConfigId(repositoryId uint) (d driver.Driver, repository model.Repository, err error) {
	repository, err = database.GetRepositoryById(repositoryId)
	if err != nil {
		return
	}
	d = pure.NewDriver(repository.ServiceAddress, repository.Username, repository.Secret, repository.Insecure)
	return
}

func GetRegDriverByRepositoryConfigId(repositoryId uint) (d driver.Driver, repository model.Repository, err error) {
	//
	//err = database.Conn.OrmShowObjectByIdUsingReflectBind(model.TableNameRepository, repositoryId, &repository)
	//if err != nil {
	//	return
	//}
	repository, err = database.GetRepositoryById(repositoryId)
	if err != nil {
		return
	}
	d, err = reg.NewDriver(repository.ServiceAddress, repository.Username, repository.Secret, repository.Insecure)
	return
}

type ListFileByRepositoryBody struct {
	model.Registry
	model.Cred
	model.RepositoryBasic
}

func GetDriverAutoParseRepositoryConfig(c *gin.Context) (driver reg.Driver, repository ListFileByRepositoryBody, err error) {
	switch c.ContentType() {
	case binding.MIMEJSON:
		err = c.BindJSON(&repository)
		if err != nil {
			lightweight_api.HandleStatusBadRequestError(c, err)
			return
		}
	case binding.MIMEMultipartPOSTForm:
		config, exists := c.GetPostForm("repository")
		if !exists {
			lightweight_api.HandleStatusBadRequestError(c, err)
			return
		}
		err = json.Unmarshal([]byte(config), &repository)
		if err != nil {
			lightweight_api.HandleStatusBadRequestError(c, err)
			return
		}
	default:
		err = errors.New("MediaType Unsupported")
		c.AbortWithStatus(http.StatusUnsupportedMediaType)
		return
	}
	driver, err = reg.NewDriver(repository.ServiceAddress, repository.Username, repository.Secret, repository.Insecure)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	return
}
