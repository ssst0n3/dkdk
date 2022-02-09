package repository

import (
	"dkdk/model"
	"dkdk/test/test_data"
	"dkdk/test/test_db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRegDriverByRepositoryConfigId(t *testing.T) {
	assert.NoError(t, test_db.InitRepository())
	driver, repository, err := GetRegDriverByRepositoryConfigId(test_data.RepositoryA.ID)
	assert.NoError(t, err)
	manifest, err := driver.GetManifests(model.RepositoryBasic{Name: repository.Name, Reference: repository.Reference})
	assert.NoError(t, err)
	assert.NotNil(t, manifest)
}
