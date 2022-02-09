package database

import (
	"dkdk/model"
	"dkdk/model/node"
	"dkdk/test/test_data"
	"dkdk/test/test_db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListDirectoryUnderDir(t *testing.T) {
	assert.NoError(t, test_db.InitDirectory())
	assert.NoError(t, test_db.InitNode())
	directories, err := ListDirectoryUnderDir(test_data.NodeA.ContentId)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(directories))
	assert.Equal(t, test_data.DirectoryB.ID, directories[0].ID)
	assert.Equal(t, test_data.DirectoryB.Filename, directories[0].Filename)
	assert.Equal(t, node.Directory, directories[0].Type)
	assert.Equal(t, test_data.NodeB.ID, directories[0].NodeId)
}

func TestGetDirectoryPath(t *testing.T) {
	assert.NoError(t, test_db.InitDirectory())
	assert.NoError(t, test_db.InitNode())
	t.Run("root node", func(t *testing.T) {
		// TODO
	})
	t.Run("sub node", func(t *testing.T) {
		path, err := GetDirectoryPath(test_data.DirectoryB.ID)
		assert.NoError(t, err)
		assert.Equal(t, []model.DirectoryPathResponse{
			{
				ID:       test_data.DirectoryA.ID,
				Filename: test_data.DirectoryA.Filename,
			},
			{
				ID:       test_data.DirectoryB.ID,
				Filename: test_data.DirectoryB.Filename,
			},
		}, path)
	})
}

func TestDeleteNodeByDirectoryId(t *testing.T) {
	assert.NoError(t, test_db.InitDirectory())
	assert.NoError(t, test_db.InitNode())
	n := node.Node{
		Core: node.Core{
			Type:      node.Directory,
			ContentId: test_data.DirectoryA.ID,
		},
	}
	assert.NoError(t, DB.First(&n).Error)
	assert.Equal(t, test_data.NodeA.ID, n.ID)
	assert.NoError(t, DeleteNodeByDirectoryId(test_data.DirectoryA.ID))
	assert.Error(t, DB.First(&n).Error)
}

func TestCreateDirectoryNode(t *testing.T) {
	filename := "DirectoryNew"
	t.Run("create directory under root", func(t *testing.T) {
		assert.NoError(t, test_db.MakeDirectoryEmpty())
		assert.NoError(t, test_db.MakeNodeEmpty())
		directoryNode, err := CreateDirectoryNode(filename, 0)
		assert.NoError(t, err)
		assert.Equal(t, uint(0), directoryNode.Parent)
		directories, err := ListDirectoryUnderDir(0)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(directories))
		assert.Equal(t, directories[0], directoryNode)
		assert.Equal(t, filename, directoryNode.Filename)
	})
	t.Run("create directory under directory", func(t *testing.T) {
		assert.NoError(t, test_db.InitDirectory())
		assert.NoError(t, test_db.InitNode())
		directoryNode, err := CreateDirectoryNode(filename, test_data.DirectoryB.ID)
		assert.NoError(t, err)
		directories, err := ListDirectoryUnderDir(test_data.DirectoryB.ID)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(directories))
		assert.Equal(t, directories[0], directoryNode)
		assert.Equal(t, filename, directoryNode.Filename)
	})
}
