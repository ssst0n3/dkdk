package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewDkdkClientForTest() *DkdkClient {
	return NewDkdkClient("127.0.0.1:14000", "admin", "admin", true)
}

func TestDkdkClient_RepositoryList(t *testing.T) {
	client := NewDkdkClientForTest()
	assert.NoError(t, client.RepositoryList())
}
