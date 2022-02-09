package node

import (
	"dkdk/model/util"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Type uint

const (
	Directory Type = iota
	File
)

type MoveBody struct {
	Parent uint `json:"parent"`
}

type Core struct {
	Type      Type `json:"type"`
	ContentId uint `json:"content_id"`
	MoveBody
}

type Node struct {
	gorm.Model
	Core
}

var (
	SchemaNode schema.Schema
)

func init() {
	awesome_error.CheckFatal(util.InitSchema(&SchemaNode, &Node{}))
}
