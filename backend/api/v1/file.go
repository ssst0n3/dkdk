package v1

import (
	"dkdk/api/v1/repository"
	"dkdk/database"
	"dkdk/model"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/response"
	"net/http"
	"strconv"
)

const FileResourceName = "file"

var FileResource = lightweight_api.NewResource(
	FileResourceName, model.SchemaFile.Table, model.File{}, "",
)

func ListFileUnderDir(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	if id != 0 {
		err = DirectoryResource.MustResourceExistsById(c, uint(id))
		if err != nil {
			return
		}
	}

	files, err := database.ListFileUnderDir(uint(id))
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, files)
}

func UploadRepositoryFileToDirectory(c *gin.Context) {
	userId, err := lightweight_api.GetUserId(c)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	id, err := repository.Resource.MustResourceExistsByIdAutoParseParam(c)
	if err != nil {
		return
	}
	var body model.UploadRepositoryFileBody
	err = c.BindJSON(&body)
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	fileId, err := database.UploadRepositoryFile(uint(id), userId, body, "")
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	response.CreateSuccess200(c, fileId, "upload repository file to net disk success")
}
