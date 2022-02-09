package database

import (
	"dkdk/test/test_data"
	"dkdk/test/test_db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAdmin(t *testing.T) {
	assert.NoError(t, test_db.InitUser())
	t.Run("admin", func(t *testing.T) {
		isAdmin, err := IsAdmin(test_data.UserAdmin.ID)
		assert.NoError(t, err)
		assert.Equal(t, test_data.UserAdmin.IsAdmin, isAdmin)
	})
	t.Run("not admin", func(t *testing.T) {
		isAdmin, err := IsAdmin(test_data.UserNormal.ID)
		assert.NoError(t, err)
		assert.Equal(t, test_data.UserNormal.IsAdmin, isAdmin)
	})
}
