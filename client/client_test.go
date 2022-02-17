package client

import (
	"github.com/ssst0n3/dkdk/model"
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

func TestDkdkClient_BatchTaskCreate(t *testing.T) {
	client := NewDkdkClientForTest()
	tasksToCreate := []model.TaskCore{
		{
			UserId: 9999,
			FileItem: model.FileItem{
				FileIdentifier: model.FileIdentifier{
					Filename: "test",
				},
			},
		},
	}
	assert.NoError(t, client.BatchTaskCreate(tasksToCreate))
	// TODO: List tasks by api
}
