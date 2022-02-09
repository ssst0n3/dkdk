package database

import (
	"errors"
	"fmt"
	"github.com/opencontainers/go-digest"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/slice"
	"github.com/ssst0n3/dkdk/model"
	"github.com/ssst0n3/dkdk/model/node"
	"gorm.io/gorm"
	"time"
)

func ListFileUnderDir(dirId uint) (files []model.FileNode, err error) {
	files = []model.FileNode{}
	joinQuery, selectQuery := JoinNodeUnderDir(model.SchemaFile.Table)
	DB.Model(&model.File{}).Select(selectQuery).Joins(joinQuery, node.File, dirId).Scan(&files)
	return
}

func DigestExists(digests []digest.Digest) (exists map[digest.Digest]string, err error) {
	exists = make(map[digest.Digest]string)
	var file []model.File
	digestName := model.SchemaFile.FieldsByName["Digest"].DBName
	query := fmt.Sprintf("%s in ?", digestName)
	err = DB.Model(&file).Where(query, digests).Find(&file).Error
	awesome_error.CheckErr(err)
	for _, f := range file {
		exists[f.Digest] = f.Filename
	}
	return
}

func CheckFileAlreadyUnderDir(file model.File, directoryId uint) (result bool, id uint, err error) {
	err = DB.Model(&model.File{}).Where(&file).First(&file).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
			return
		}
		awesome_error.CheckErr(err)
		return
	}
	fileNode := node.Node{
		Core: node.Core{
			Type:      node.File,
			ContentId: file.ID,
			MoveBody: node.MoveBody{
				Parent: directoryId,
			},
		},
	}
	err = DB.Model(&node.Node{}).Where(&fileNode).First(&fileNode).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
			return
		}
		awesome_error.CheckErr(err)
		return
	}
	result = true
	id = fileNode.ID
	return
}

func UploadRepositoryFile(repositoryId uint, userId uint, body model.UploadRepositoryFileBody, password string) (id uint, err error) {
	if password == "" {
		password, err = QueryPasswordByDigest(userId, repositoryId, body.Digest)
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
	}
	file := model.File{
		FileItem:        body.FileItem,
		UserId:          userId,
		RepositoryId:    repositoryId,
		ArchivePassword: password,
	}

	exists, id, err := CheckFileAlreadyUnderDir(file, body.DirectoryId)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	if exists {
		return
	}
	file.LastModified = time.Now()
	err = DB.Create(&file).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	fileNode := node.Node{
		Core: node.Core{
			Type:      node.File,
			ContentId: file.ID,
			MoveBody: node.MoveBody{
				Parent: body.DirectoryId,
			},
		},
	}
	err = DB.Create(&fileNode).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	id = fileNode.ID
	return
}

func getNodeParent(child node.Node) (parent node.Node, err error) {
	err = DB.Model(&node.Node{}).First(&parent, child.Parent).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func CheckNodeCycle(n node.Node) (result bool, err error) {
	var ids []uint
	for n.Parent != 0 {
		ids = append(ids, n.ID)
		n, err = getNodeParent(n)
		if err != nil {
			return
		}
		if slice.In(n.ID, ids) {
			result = true
			return
		}
	}
	return
}

func FindFilesByDigest(dsg digest.Digest) (files []model.File, err error) {
	err = DB.Model(&model.File{}).Where(&model.File{
		FileItem: model.FileItem{
			FileIdentifier: model.FileIdentifier{
				Digest: dsg,
			},
		},
	}).Find(&files).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func UpdateFileDownloadCount(dsg digest.Digest) (err error) {
	files, err := FindFilesByDigest(dsg)
	if err != nil {
		return
	}
	for _, file := range files {
		count := file.DownloadCount
		if count < 65536 {
			count += 1
		}
		err = DB.Model(&model.File{}).Where(&file).Update(model.SchemaFile.FieldsByName["DownloadCount"].DBName, count).Error
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
	}
	return
}
