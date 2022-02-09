package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {
	token, err := Login("http://127.0.0.1:14000/api/v1/auth", "admin", "admin")
	assert.NoError(t, err)
	assert.True(t, len(token) > 0)
}
