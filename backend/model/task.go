package model

import (
	"dkdk/model/util"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	TaskTypeOfflineDownload uint = iota
)

const (
	TaskStatusWaiting uint = iota
	TaskStatusDownloading
	TaskStatusDownloaded
	TaskStatusArchived
	TaskStatusUploading
	TaskStatusUploaded
	TaskStatusDone
)

type TaskCore struct {
	UserId          uint   `json:"user_id"`
	RepositoryId    uint   `json:"repository_id"`
	Type            uint   `json:"type"`
	OriginUrl       string `json:"origin_url"`
	Status          uint   `json:"status"`
	ArchivePassword string `json:"archive_password"`
	FilenameInDkdk  string `json:"filename_in_dkdk"`
	DirectoryInDkdk uint   `json:"directory_in_dkdk"`
	FileItem
}

type Task struct {
	gorm.Model
	TaskCore
}

type TaskWithId struct {
	Id uint `json:"id"`
	Task
}

var (
	SchemaTask schema.Schema
)

func init() {
	awesome_error.CheckFatal(util.InitSchema(&SchemaTask, &Task{}))
}
