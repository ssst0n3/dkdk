package database

import (
	"github.com/opencontainers/go-digest"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/dkdk/model"
	"gorm.io/gorm"
)

func Paginate(page, pageSize uint) func(db *gorm.DB) *gorm.DB {
	if page == 0 {
		page = 1
	}
	offset := int((page - 1) * pageSize)
	return func(db *gorm.DB) *gorm.DB {
		if pageSize == 0 {
			return db
		}
		return db.Offset(offset).Limit(int(pageSize))
	}
}
func ListTaskByUserId(id, page, size uint) (tasks []model.Task, err error) {
	err = DB.Scopes(Paginate(page, size)).Limit(int(size)).Where(model.Task{
		TaskCore: model.TaskCore{
			UserId: id,
		},
	}).Order("id desc").Find(&tasks).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

/*
BatchCreateTask
Make sure user exists by you self.
*/
func BatchCreateTask(tasks []model.Task, userId uint) (err error) {
	for i := range tasks {
		tasks[i].UserId = userId
	}
	err = DB.Create(&tasks).Error
	awesome_error.CheckErr(err)
	return
}

func UpdateTask(task model.Task) (err error) {
	err = DB.Updates(&task).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func UpdateTaskStatus(id uint, status uint) (err error) {
	task := model.Task{
		Model: gorm.Model{
			ID: id,
		},
		TaskCore: model.TaskCore{
			Status: status,
		},
	}
	err = DB.Updates(&task).Error
	awesome_error.CheckErr(err)
	return
}

func UpdateTaskArchivePassword(id uint, password string) (err error) {
	task := model.Task{
		Model: gorm.Model{
			ID: id,
		},
		TaskCore: model.TaskCore{
			ArchivePassword: password,
		},
	}
	err = DB.Updates(&task).Error
	awesome_error.CheckErr(err)
	return
}

func QueryPasswordByDigest(userId uint, repositoryId uint, digest digest.Digest) (password string, err error) {
	task := model.Task{
		TaskCore: model.TaskCore{
			UserId:       userId,
			RepositoryId: repositoryId,
			Status:       uint(model.TaskStatusDone),
			FileItem: model.FileItem{
				FileIdentifier: model.FileIdentifier{
					Digest: digest,
				},
			},
		},
	}
	err = DB.Model(&task).Where(&task).Select(model.SchemaTask.FieldsByName["ArchivePassword"].DBName).First(&password).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		} else {
			awesome_error.CheckErr(err)
			return
		}
	}
	return
}
