package test_data

import (
	"github.com/opencontainers/go-digest"
	"github.com/ssst0n3/dkdk/model"
	"gorm.io/gorm"
)

var (
	Task1 = model.Task{
		Model: gorm.Model{ID: 1},
		TaskCore: model.TaskCore{
			UserId:          UserAdmin.ID,
			RepositoryId:    RepositoryA.ID,
			Status:          model.TaskStatusDone,
			ArchivePassword: "password",
			FileItem: model.FileItem{
				FileIdentifier: model.FileIdentifier{
					Digest: digest.FromString("file"),
				},
			},
		},
	}
	Task2 = model.Task{
		Model: gorm.Model{ID: 2},
		TaskCore: model.TaskCore{
			UserId:       UserNormal.ID,
			RepositoryId: RepositoryA.ID,
		},
	}
	Task3 = model.Task{
		Model: gorm.Model{ID: 3},
		TaskCore: model.TaskCore{
			UserId:       UserAdmin.ID,
			RepositoryId: RepositoryA.ID,
		},
	}
	Task4 = model.Task{
		Model: gorm.Model{ID: 4},
		TaskCore: model.TaskCore{
			UserId:       UserAdmin.ID,
			RepositoryId: RepositoryA.ID,
		},
	}
	Tasks = []model.Task{
		Task1,
		Task2,
		Task3,
		Task4,
	}
)
