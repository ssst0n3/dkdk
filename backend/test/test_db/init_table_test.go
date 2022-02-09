package test_db

import (
	"dkdk/test/test_data"
	"github.com/ssst0n3/awesome_libs/cipher"
	"github.com/ssst0n3/lightweight_api/example/resource/user/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitRepository(t *testing.T) {
	assert.NoError(t, InitRepository())
}

func TestInitNode(t *testing.T) {

}

func TestInitUser(t *testing.T) {
	cipher.Init()
	assert.NoError(t, InitUser())
	var users []model.User
	assert.NoError(t, db.Find(&users).Error)
	password, err := cipher.CommonCipher.Decrypt(users[0].Password)
	assert.NoError(t, err)
	assert.Equal(t, test_data.UserAdmin.Password, string(password))
}

func TestInitAllTables(t *testing.T) {
	cipher.Init()
	assert.NoError(t, InitAllTables())
}
