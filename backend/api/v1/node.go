package v1

import (
	"dkdk/database"
	"dkdk/model/node"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/response"
)

const NodeResourceName = "node"

var NodeResource = lightweight_api.NewResource(
	NodeResourceName,
	node.SchemaNode.Table,
	node.Node{},
	"",
)

func BatchDeleteNode(c *gin.Context) {
	var ids []uint
	err := c.BindJSON(&ids)
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	err = database.BatchDeleteNode(ids)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	response.DeleteSuccess200(c)
}

func Move(c *gin.Context) {
	id, err := NodeResource.MustResourceExistsByIdAutoParseParam(c)
	if err != nil {
		return
	}
	var n node.MoveBody
	if err := c.ShouldBindJSON(&n); err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	err = database.DB.Table(node.SchemaNode.Table).Where("id", id).Save(&n).Error
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	return
}
