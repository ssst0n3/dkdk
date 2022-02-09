package test_data

import (
	"dkdk/model"
	"gorm.io/gorm"
)

var (
	RepositoryA = model.Repository{
		Model: gorm.Model{
			ID: 1,
		},
		UserId:  UserAdmin.ID,
		Type:    model.DockerRegistry,
		Private: false,
		Registry: model.Registry{
			ServiceAddress: "172.17.0.1:14005",
			Insecure:       true,
		},
		Cred: model.Cred{},
		RepositoryBasic: model.RepositoryBasic{
			Name:      "dkdk/hello-world",
			Reference: "v2",
		},
	}
	Repositories = []model.Repository{
		RepositoryA,
	}
)
