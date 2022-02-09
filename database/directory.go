package database

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/dkdk/model"
	"github.com/ssst0n3/dkdk/model/node"
	"gorm.io/gorm"
)

func ListDirectoryUnderDir(dirId uint) (directories []model.DirectoryNode, err error) {
	directories = []model.DirectoryNode{}
	joinQuery, selectQuery := JoinNodeUnderDir(model.SchemaDirectory.Table)
	DB.Model(&model.Directory{}).Select(selectQuery).Joins(joinQuery, node.Directory, dirId).Scan(&directories)
	return
}

func GetDirectoryPath(id uint) (path []model.DirectoryPathResponse, err error) {
	if id == 0 {
		path = []model.DirectoryPathResponse{}
		return
	}
	dir := model.Directory{}
	err = DB.First(&dir, id).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	dirNode := node.Node{
		Core: node.Core{
			Type:      node.Directory,
			ContentId: dir.ID,
		},
	}
	err = DB.Where(&dirNode).First(&dirNode).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	parentPath, err := GetDirectoryPath(dirNode.Parent)
	if err != nil {
		return
	}
	path = append(parentPath, model.DirectoryPathResponse{
		Filename: dir.Filename,
		ID:       dir.ID,
	})
	return
}

func DeleteNodeByDirectoryId(directoryId uint) (err error) {
	err = DB.Unscoped().Where(&node.Node{
		Core: node.Core{
			Type:      node.Directory,
			ContentId: directoryId,
		},
	}).Delete(&node.Node{}).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func CreateDirectoryNode(filename string, parentNodeId uint) (directoryNode model.DirectoryNode, err error) {
	directory := model.Directory{Filename: filename}
	err = DB.Create(&directory).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	n := node.Node{
		Core: node.Core{
			Type:      node.Directory,
			ContentId: directory.ID,
			MoveBody: node.MoveBody{
				Parent: parentNodeId,
			},
		},
	}
	err = DB.Create(&n).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	joinQuery, selectQuery := JoinNodeUnderDir(model.SchemaDirectory.Table)
	err = DB.Model(&model.Directory{}).Select(selectQuery).Joins(joinQuery, node.Directory, parentNodeId).Where(
		&model.Directory{Model: gorm.Model{ID: directory.ID}},
	).First(&directoryNode).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
