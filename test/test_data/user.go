package test_data

import (
	"github.com/ssst0n3/lightweight_api/example/resource/user/model"
	"gorm.io/gorm"
)

var (
	UserAdmin = model.User{
		Model: gorm.Model{ID: 1},
		CreateUserBody: model.CreateUserBody{
			UpdateBasicBody: model.UpdateBasicBody{
				Username: "admin",
				IsAdmin:  true,
			},
			UpdatePasswordBody: model.UpdatePasswordBody{
				Password: "admin",
			},
		},
	}
	UserNormal = model.User{
		Model: gorm.Model{ID: 2},
		CreateUserBody: model.CreateUserBody{
			UpdateBasicBody: model.UpdateBasicBody{
				Username: "user",
				IsAdmin:  false,
			},
			UpdatePasswordBody: model.UpdatePasswordBody{
				Password: "user",
			},
		},
	}
	Users = []model.User{
		UserAdmin,
		UserNormal,
	}
)
