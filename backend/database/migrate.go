package database

import (
	modelDkDk "dkdk/model"
	"dkdk/model/node"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/awesome_reflect"
	modelKV "github.com/ssst0n3/lightweight_api/example/resource/kv_config/model"
	modelUser "github.com/ssst0n3/lightweight_api/example/resource/user/model"
)

func migrate(m interface{}) (err error) {
	awesome_reflect.MustPointer(m)
	err = DB.AutoMigrate(m)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func MigrateTables(models ...interface{}) (err error) {
	for _, m := range models {
		err = migrate(m)
		if err != nil {
			return err
		}
	}
	return
}

func Migrate() (err error) {
	return MigrateTables(
		&modelDkDk.Repository{},
		&modelDkDk.Task{},
		&modelUser.User{},
		&node.Node{},
		&modelDkDk.File{},
		&modelDkDk.Directory{},
		&modelKV.Config{},
	)
}
