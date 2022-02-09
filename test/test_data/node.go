package test_data

import (
	"github.com/ssst0n3/dkdk/model/node"
	"gorm.io/gorm"
)

var (
	NodeA = node.Node{
		Model: gorm.Model{
			ID: 1,
		},
		Core: node.Core{
			Type:      node.Directory,
			ContentId: DirectoryA.ID,
		},
	}
	NodeB = node.Node{
		Model: gorm.Model{
			ID: 2,
		},
		Core: node.Core{
			Type:      node.Directory,
			ContentId: DirectoryB.ID,
			MoveBody: node.MoveBody{
				Parent: NodeA.ID,
			},
		},
	}
	NodeC = node.Node{
		Model: gorm.Model{
			ID: 3,
		},
		Core: node.Core{
			Type:      node.File,
			ContentId: FileA.ID,
		},
	}
	NodeD = node.Node{
		Model: gorm.Model{
			ID: 4,
		},
		Core: node.Core{
			Type:      node.File,
			ContentId: FileB.ID,
			MoveBody: node.MoveBody{
				Parent: NodeA.ID,
			},
		},
	}
	NodeE = node.Node{
		Model: gorm.Model{
			ID: 5,
		},
		Core: node.Core{
			MoveBody: node.MoveBody{
				Parent: NodeCycle,
			},
		},
	}
	NodeCycle = uint(6)
	NodeF     = node.Node{
		Model: gorm.Model{
			ID: NodeCycle,
		},
		Core: node.Core{
			MoveBody: node.MoveBody{
				Parent: NodeE.ID,
			},
		},
	}
	Nodes = []node.Node{
		NodeA,
		NodeB,
		NodeC,
		NodeD,
		NodeE,
		NodeF,
	}
)
