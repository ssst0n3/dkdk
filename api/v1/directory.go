package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/dkdk/database"
	"github.com/ssst0n3/dkdk/model"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/response"
	"net/http"
	"strconv"
)

const DirectoryResourceName = "directory"

var DirectoryResource = lightweight_api.NewResource(
	DirectoryResourceName, model.SchemaDirectory.Table, model.Directory{}, "",
)

func ListDirectoryUnderDir(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	if id != 0 {
		err = DirectoryResource.MustResourceExistsByGuid(c, model.SchemaDirectory.FieldsByName["ID"].DBName, id)
		if err != nil {
			return
		}
	}
	directories, err := database.ListDirectoryUnderDir(uint(id))
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, directories)
}

func GetDirectoryPath(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	err = DirectoryResource.MustResourceExistsById(c, uint(id))
	if err != nil {
		return
	}
	path, err := database.GetDirectoryPath(uint(id))
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, path)
}

func DeleteNodeByDirectoryId(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	err = DirectoryResource.MustResourceExistsById(c, uint(id))
	if err != nil {
		return
	}
	err = database.DeleteNodeByDirectoryId(uint(id))
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	response.DeleteSuccess200(c)
}

func CreateDirectoryNode(c *gin.Context) {
	var body model.CreateDirectoryBody
	err := c.BindJSON(&body)
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	directoryNode, err := database.CreateDirectoryNode(body.Filename, body.ParentNodeId)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, directoryNode)
}
