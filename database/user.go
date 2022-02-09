package database

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/lightweight_api/example/resource/user/model"
	"gorm.io/gorm"
)

func IsAdmin(userId uint) (isAdmin bool, err error) {
	user := model.User{
		Model: gorm.Model{
			ID: userId,
		},
	}
	err = DB.First(&user).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	isAdmin = user.IsAdmin
	return
}
