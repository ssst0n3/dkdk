package database

import (
	"github.com/opencontainers/go-digest"
	"github.com/ssst0n3/dkdk/model"
	"github.com/ssst0n3/dkdk/model/node"
	"github.com/ssst0n3/dkdk/test/test_data"
	"github.com/ssst0n3/dkdk/test/test_db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListFileUnderDir(t *testing.T) {
	assert.NoError(t, test_db.InitFile())
	assert.NoError(t, test_db.InitNode())
	files, err := ListFileUnderDir(test_data.NodeA.ContentId)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(files))
	assert.Equal(t, test_data.FileB.ID, files[0].ID)
	assert.Equal(t, test_data.FileB.Filename, files[0].Filename)
	assert.Equal(t, node.File, files[0].Type)
	assert.Equal(t, test_data.NodeD.ID, files[0].NodeId)
}

func TestDigestExists(t *testing.T) {
	assert.NoError(t, test_db.InitFile())
	exists, err := DigestExists([]digest.Digest{test_data.FileA.Digest, digest.FromString("not exists")})
	assert.NoError(t, err)
	assert.Equal(t, map[digest.Digest]string{
		test_data.FileA.Digest: test_data.FileA.Filename,
	}, exists)
}

func TestCheckFileAlreadyUnderDir(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		assert.NoError(t, test_db.InitAllTables())
		file := model.File{
			FileItem:        test_data.FileA.FileItem,
			UserId:          test_data.FileA.UserId,
			RepositoryId:    test_data.FileA.RepositoryId,
			ArchivePassword: test_data.FileA.ArchivePassword,
		}
		result, id, err := CheckFileAlreadyUnderDir(file, test_data.NodeA.Parent)
		assert.NoError(t, err)
		assert.Equal(t, true, result)
		assert.Equal(t, test_data.NodeC.ID, id)
	})
	t.Run("not exists", func(t *testing.T) {
		assert.NoError(t, test_db.MakeFileEmpty())
		file := model.File{
			FileItem:        test_data.FileA.FileItem,
			UserId:          test_data.FileA.UserId,
			RepositoryId:    test_data.FileA.RepositoryId,
			ArchivePassword: test_data.FileA.ArchivePassword,
		}
		result, _, err := CheckFileAlreadyUnderDir(file, test_data.DirectoryA.ID)
		assert.NoError(t, err)
		assert.Equal(t, false, result)
	})
}

func TestUploadRepositoryFile(t *testing.T) {
	assert.NoError(t, test_db.InitUser())
	assert.NoError(t, test_db.InitRepository())
	assert.NoError(t, test_db.InitDirectory())
	assert.NoError(t, test_db.MakeFileEmpty())
	assert.NoError(t, test_db.InitNode())
	assert.NoError(t, test_db.InitTask())
	body := model.UploadRepositoryFileBody{
		FileItem: model.FileItem{
			FileIdentifier: model.FileIdentifier{
				Filename: "new_upload_file",
				Digest:   digest.FromString("new_upload_file"),
			},
			Size: 1000,
		},
		DirectoryId: test_data.NodeA.ContentId,
	}
	id, err := UploadRepositoryFile(test_data.RepositoryA.ID, test_data.UserAdmin.ID, body, "")
	assert.NoError(t, err)
	var fileNode node.Node
	assert.NoError(t, DB.First(&fileNode, id).Error)
	files, err := ListFileUnderDir(body.DirectoryId)
	assert.NoError(t, err)
	assert.Equal(t, fileNode.ContentId, files[0].ID)
}

func TestCheckNodeCycle(t *testing.T) {
	assert.NoError(t, test_db.InitAllTables())
	t.Run("no cycle", func(t *testing.T) {
		cycle, err := CheckNodeCycle(test_data.NodeD)
		assert.NoError(t, err)
		assert.False(t, cycle)
	})
	t.Run("cycle", func(t *testing.T) {
		cycle, err := CheckNodeCycle(test_data.NodeE)
		assert.NoError(t, err)
		assert.True(t, cycle)
	})
}

func TestFindFilesByDigest(t *testing.T) {
	assert.NoError(t, test_db.InitFile())
	t.Run("not exists", func(t *testing.T) {
		files, err := FindFilesByDigest("not_exists")
		assert.NoError(t, err)
		assert.Equal(t, []model.File{}, files)
	})
	t.Run("success", func(t *testing.T) {
		files, err := FindFilesByDigest(test_data.FileA.Digest)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(files))
		assert.Equal(t, test_data.FileA.ID, files[0].ID)
	})
}

func TestUpdateFileDownloadCount(t *testing.T) {
	assert.NoError(t, test_db.InitFile())
	t.Run("not exists", func(t *testing.T) {
		assert.NoError(t, UpdateFileDownloadCount("not_exists"))
		var file model.File
		assert.NoError(t, DB.Model(&model.File{}).First(&file, test_data.FileA.ID).Error)
		assert.Equal(t, uint(0), file.DownloadCount)
	})
	t.Run("success", func(t *testing.T) {
		assert.NoError(t, UpdateFileDownloadCount(test_data.FileA.Digest))
		var file model.File
		assert.NoError(t, DB.Model(&model.File{}).First(&file, test_data.FileA.ID).Error)
		assert.Equal(t, uint(1), file.DownloadCount)
	})
}
