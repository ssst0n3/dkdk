package database

import (
	"dkdk/model/node"
	"github.com/ssst0n3/awesome_libs"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

func JoinNodeUnderDir(fileTable string) (joinQuery, selectQuery string) {
	joinQuery = awesome_libs.Format(
		"JOIN {.node_table} "+
			"ON {.node_table}.{.content_id}={.file_table}.id "+
			"AND {.node_table}.{.type}=? "+
			"AND {.node_table}.{.parent}=?",
		awesome_libs.Dict{
			"node_table": node.SchemaNode.Table,
			"content_id": node.SchemaNode.FieldsByName["ContentId"].DBName,
			"file_table": fileTable,
			"type":       node.SchemaNode.FieldsByName["Type"].DBName,
			"parent":     node.SchemaNode.FieldsByName["Parent"].DBName,
		})
	selectQuery = awesome_libs.Format(
		"{.node_table}.*, {.node_table}.id as {.node_id}, {.file_table}.*",
		awesome_libs.Dict{
			"node_table": node.SchemaNode.Table,
			"file_table": fileTable,
			"node_id":    "node_id",
		},
	)
	return
}

func BatchDeleteNode(ids []uint) (err error) {
	err = DB.Unscoped().Delete(&node.Node{}, ids).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
