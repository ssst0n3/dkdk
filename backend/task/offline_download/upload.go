package offline_download

import (
	"dkdk/database"
	"dkdk/model"
	"dkdk/registry/v2/driver/reg"
	"github.com/opencontainers/go-digest"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"gorm.io/gorm"
	"os"
)

func Upload(taskId uint, repository model.Repository, path string, filename string) (dgs digest.Digest, size int64, err error) {
	log.Logger.Info()
	driver, err := reg.NewDriver(repository.ServiceAddress, repository.Username, repository.Secret, repository.Insecure)
	if err != nil {
		return
	}
	content, err := os.Open(path)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	dgs, size, err = Info(path)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	task := model.Task{
		Model: gorm.Model{
			ID: taskId,
		},
		TaskCore: model.TaskCore{
			FileItem: model.FileItem{
				FileIdentifier: model.FileIdentifier{
					Digest: dgs,
				},
				Size: size,
			},
		},
	}
	err = database.UpdateTask(task)
	if err != nil {
		return
	}
	err = driver.Upload(
		model.RepositoryBasic{Name: repository.Name, Reference: repository.Reference},
		filename, content, dgs, size,
	)
	return
}
