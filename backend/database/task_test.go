package database

import (
	"dkdk/model"
	"dkdk/test/test_data"
	"dkdk/test/test_db"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestListTaskByUserId(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		assert.NoError(t, test_db.InitTask())
		tasks, err := ListTaskByUserId(test_data.Task1.UserId, 0, 0)
		assert.NoError(t, err)
		assert.Equal(t, len(test_data.Tasks)-1, len(tasks))
		assert.Equal(t, test_data.Task1.RepositoryId, tasks[0].RepositoryId)
	})
	t.Run("page size 2", func(t *testing.T) {
		assert.NoError(t, test_db.InitTask())
		tasks, err := ListTaskByUserId(test_data.Task1.UserId, 0, 2)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(tasks))
		assert.Equal(t, test_data.Task1.RepositoryId, tasks[0].RepositoryId)
	})
	t.Run("page size 100", func(t *testing.T) {
		assert.NoError(t, test_db.InitTask())
		tasks, err := ListTaskByUserId(test_data.Task1.UserId, 0, 100)
		assert.NoError(t, err)
		assert.Equal(t, len(test_data.Tasks)-1, len(tasks))
		assert.Equal(t, test_data.Task1.RepositoryId, tasks[0].RepositoryId)
	})
}

func TestBatchCreateTask(t *testing.T) {
	assert.NoError(t, test_db.MakeTaskEmpty())
	userNotExists := uint(9999)
	assert.NoError(t, BatchCreateTask(test_data.Tasks, userNotExists))
	tasks, err := ListTaskByUserId(userNotExists, 0, 0)
	assert.NoError(t, err)
	assert.Equal(t, len(test_data.Tasks), len(tasks))
	assert.Equal(t, userNotExists, tasks[0].UserId)
}

func TestUpdateTaskStatus(t *testing.T) {
	assert.NoError(t, test_db.InitTask())
	assert.NotEqual(t, model.TaskStatusDone, test_data.Task1.ID)
	assert.NoError(t, UpdateTaskStatus(test_data.Task1.ID, model.TaskStatusDone))
	assert.NoError(t, DB.First(&test_data.Task1).Error)
	assert.Equal(t, model.TaskStatusDone, test_data.Task1.Status)
}

func TestUpdateTaskArchivePassword(t *testing.T) {
	assert.NoError(t, test_db.InitTask())
	password := time.Now().String()
	assert.NoError(t, UpdateTaskArchivePassword(test_data.Task1.ID, password))
	assert.NoError(t, DB.First(&test_data.Task1).Error)
	assert.Equal(t, password, test_data.Task1.ArchivePassword)
}

func TestQueryPasswordByDigest(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.NoError(t, test_db.MakeTaskEmpty())
		password, err := QueryPasswordByDigest(test_data.Task1.UserId, test_data.Task1.RepositoryId, test_data.Task1.Digest)
		assert.NoError(t, err)
		assert.Equal(t, "", password)
	})
	t.Run("not empty", func(t *testing.T) {
		assert.NoError(t, test_db.InitTask())
		password, err := QueryPasswordByDigest(test_data.Task1.UserId, test_data.Task1.RepositoryId, test_data.Task1.Digest)
		assert.NoError(t, err)
		assert.Equal(t, "password", password)
	})
	t.Run("not empty2", func(t *testing.T) {
		assert.NoError(t, test_db.InitTask())
		password, err := QueryPasswordByDigest(test_data.Task1.UserId, test_data.Task1.RepositoryId, "another digest")
		assert.NoError(t, err)
		assert.Equal(t, "", password)
	})
}
