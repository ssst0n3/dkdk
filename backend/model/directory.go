package model

import (
	"dkdk/model/node"
	"dkdk/model/util"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/*
Directory
DirName = Node.Filename
*/
type Directory struct {
	gorm.Model
	Filename string `json:"filename"` // be same with File
}

type DirectoryNode struct {
	Directory
	node.Core
	NodeId uint `json:"node_id"`
}

type CreateDirectoryBody struct {
	Filename     string `json:"filename"`
	ParentNodeId uint   `json:"parent_node_id"`
}

type DirectoryPathResponse struct {
	ID       uint   `json:"id"`
	Filename string `json:"filename"`
}

var (
	SchemaDirectory schema.Schema
)

func init() {
	awesome_error.CheckFatal(util.InitSchema(&SchemaDirectory, &Directory{}))
}
