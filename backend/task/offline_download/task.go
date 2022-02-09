package offline_download

import (
	"dkdk/database"
	"dkdk/model"
	"fmt"
	"github.com/google/uuid"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"os"
)

var ChanTask = make(chan model.Task, 1)

func Run() {
	log.Logger.Info()
	for {
		task := <-ChanTask
		_ = OfflineDownload(task)
	}
}

func OfflineDownload(task model.Task) (err error) {
	log.Logger.Infof("start task: %s", task.OriginUrl)
	path := fmt.Sprintf("/tmp/%s", uuid.New().String())
	err = database.UpdateTaskStatus(task.ID, model.TaskStatusDownloading)
	if err != nil {
		return
	}
	err = Download(task.OriginUrl, path)
	if err != nil {
		return
	}
	err = database.UpdateTaskStatus(task.ID, model.TaskStatusDownloaded)
	if err != nil {
		return
	}
	password := uuid.New().String()
	archivePath, err := Archive(path, task.Filename, password)
	if err != nil {
		return
	}
	err = database.UpdateTaskArchivePassword(task.ID, password)
	if err != nil {
		return
	}
	err = database.UpdateTaskStatus(task.ID, model.TaskStatusArchived)
	if err != nil {
		return
	}
	repositoryConfig, err := database.GetRepositoryById(task.RepositoryId)
	if err != nil {
		return
	}
	dgs, size, err := Upload(task.ID, repositoryConfig, archivePath, task.Filename)
	if err != nil {
		return
	}
	err = database.UpdateTaskStatus(task.ID, model.TaskStatusUploaded)
	if err != nil {
		return
	}
	err = os.Remove(path)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	err = os.Remove(archivePath)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	if len(task.FilenameInDkdk) > 0 {
		body := model.UploadRepositoryFileBody{
			FileItem: model.FileItem{
				FileIdentifier: model.FileIdentifier{
					Filename: task.FilenameInDkdk,
					Digest:   dgs,
				},
				Size: size,
			},
			DirectoryId: task.DirectoryInDkdk,
		}
		_, err = database.UploadRepositoryFile(task.RepositoryId, task.UserId, body, password)
		if err != nil {
			return
		}
	}
	err = database.UpdateTaskStatus(task.ID, model.TaskStatusDone)
	if err != nil {
		return
	}
	return
}
