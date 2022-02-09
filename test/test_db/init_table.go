package test_db

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/dkdk/model"
	"github.com/ssst0n3/dkdk/model/node"
	"github.com/ssst0n3/dkdk/test/test_common"
	"github.com/ssst0n3/dkdk/test/test_data"
	"github.com/ssst0n3/lightweight_api/example/resource/user"
	model2 "github.com/ssst0n3/lightweight_api/example/resource/user/model"
)

func MakeNodeEmpty() (err error) {
	return test_common.MakeTableEmpty(db, &node.Node{})
}

func InitNode() (err error) {
	return test_common.InitTable(db, &node.Node{}, &test_data.Nodes)
}

func MakeFileEmpty() (err error) {
	return test_common.MakeTableEmpty(db, &model.File{})
}

func InitFile() (err error) {
	return test_common.InitTable(db, &model.File{}, &test_data.Files)
}

func MakeDirectoryEmpty() (err error) {
	return test_common.MakeTableEmpty(db, &model.Directory{})
}

func InitDirectory() (err error) {
	return test_common.InitTable(db, &model.Directory{}, &test_data.Directories)
}

func InitRepository() (err error) {
	return test_common.InitTable(db, &model.Repository{}, &test_data.Repositories)
}

func InitUser() (err error) {
	var users []model2.User
	for _, u := range test_data.Users {
		err = user.EncryptUser(&u)
		if err != nil {
			awesome_error.CheckErr(err)
			return err
		}
		users = append(users, u)
	}

	return test_common.InitTable(db, &model2.User{}, &users)
}

func MakeTaskEmpty() (err error) {
	return test_common.MakeTableEmpty(db, &model.Task{})
}

func InitTask() (err error) {
	return test_common.InitTable(db, &model.Task{}, &test_data.Tasks)
}

func InitAllTables() (err error) {
	err = InitUser()
	if err != nil {
		return
	}
	err = InitNode()
	if err != nil {
		return
	}
	err = InitTask()
	if err != nil {
		return
	}
	err = InitDirectory()
	if err != nil {
		return
	}
	err = InitFile()
	if err != nil {
		return
	}
	err = InitRepository()
	if err != nil {
		return
	}
	return
}
