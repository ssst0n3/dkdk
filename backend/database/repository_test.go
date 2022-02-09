package database

import (
	"dkdk/test/test_data"
	"dkdk/test/test_db"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListRepository(t *testing.T) {
	assert.NoError(t, test_db.InitRepository())
	result, err := ListRepositoryForResponse(test_data.RepositoryA.UserId)
	assert.NoError(t, err)
	log.Logger.Info(result)
	assert.Equal(t, test_data.RepositoryA.ID, result[0].Id)
	assert.Equal(t, test_data.RepositoryA.RepositoryBasic, result[0].RepositoryBasic)
}

func TestGetRepositoryById(t *testing.T) {
	assert.NoError(t, test_db.InitRepository())
	repository, err := GetRepositoryById(test_data.RepositoryA.ID)
	assert.NoError(t, err)
	assert.Equal(t, test_data.RepositoryA.ID, repository.ID)
	assert.Equal(t, test_data.RepositoryA.UserId, repository.UserId)
}

func TestListRepositoryByUserId(t *testing.T) {
	assert.NoError(t, test_db.InitRepository())
	repository, err := ListRepositoryByUserId(test_data.RepositoryA.UserId)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(repository))
	assert.Equal(t, test_data.RepositoryA.Registry, repository[0].Registry)
}

func TestCheckUserHasRepository(t *testing.T) {
	t.Run("yes", func(t *testing.T) {
		has, err := CheckUserHasRepository(test_data.RepositoryA.UserId, test_data.RepositoryA.ID)
		assert.NoError(t, err)
		assert.True(t, has)
	})
	t.Run("no", func(t *testing.T) {
		has, err := CheckUserHasRepository(test_data.RepositoryA.UserId+1, test_data.RepositoryA.ID)
		assert.NoError(t, err)
		assert.False(t, has)
	})
}
