package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/dkdk/api/v1/repository"
	"github.com/ssst0n3/dkdk/database"
	"github.com/ssst0n3/dkdk/model"
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

func CheckFilenameAlreadyExists(c *gin.Context) {
	filename := c.Query("filename")
	if filename == "" {
		lightweight_api.HandleStatusBadRequestError(c, fmt.Errorf("filename cannot be empty"))
		return
	}
	exists, err := database.CheckFilenameAlreadyExists(filename)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	resp := model.ResponseCheckFilenameAlreadyExists{Exists: exists}
	c.JSON(http.StatusOK, resp)
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
