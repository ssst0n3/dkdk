package repository

import (
	"dkdk/database"
	"dkdk/model"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api"
	"net/http"
)

const ResourceName = "repository"

var Resource = lightweight_api.NewResource(
	ResourceName, model.SchemaRepository.Table, model.Repository{}, "",
)

func ListRepository(c *gin.Context) {
	userId, err := lightweight_api.GetUserId(c)
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}

	//err = database.Conn.OrmShowObjectOnePropertyByIdByReflectBind(user.ResourceName, user.ColumnNameIsAdmin, int64(userId), &isAdmin)
	isAdmin, err := database.IsAdmin(userId)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	if isAdmin {
		userId = 0
	}
	objects, err := database.ListRepositoryForResponse(userId)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, objects)
}

func GetRepositoryFromParamId(c *gin.Context) (repository model.Repository, err error) {
	id, err := Resource.MustResourceExistsByIdAutoParseParam(c)
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	repository, err = database.GetRepositoryById(uint(id))
	return
}

func AddRepository(c *gin.Context) {
	Resource.CreateResourceTemplate(c, func(modelPtr interface{}) (err error) {
		repositoryConfig := modelPtr.(*model.Repository)
		repositoryConfig.UserId, err = lightweight_api.GetUserId(c)
		if err != nil {
			lightweight_api.HandleInternalServerError(c, err)
			return
		}
		return
	}, nil)
}

func ShowRepository(c *gin.Context) {

}

func UpdateRepository(c *gin.Context) {
	Resource.UpdateResource(c)
}
