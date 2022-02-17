package model

import (
	"github.com/opencontainers/go-digest"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/dkdk/model/node"
	"github.com/ssst0n3/dkdk/model/util"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type FileIdentifier struct {
	Filename string        `json:"filename"`
	Digest   digest.Digest `json:"digest"`
}

type FileItem struct {
	FileIdentifier
	Size int64 `json:"size"`
}

type ResponseListFileUnderRepository struct {
	FileItem
	FilenameInDkdk string `json:"filename_in_dkdk"`
	InNetDisk      bool   `json:"in_net_disk"`
}

type ResponseCheckFilenameAlreadyExists struct {
	Exists bool `json:"exists"`
}

type UploadRepositoryFileBody struct {
	FileItem
	DirectoryId uint `json:"directory_id"`
}

type File struct {
	gorm.Model
	FileItem
	UserId               uint      `json:"user_id"`
	RepositoryId         uint      `json:"repository_id"`
	FilenameInRepository string    `json:"filename_in_repository"`
	OriginUrl            string    `json:"origin_url"`
	ArchivePassword      string    `json:"archive_password"`
	LastModified         time.Time `json:"last_modified"`
	DownloadCount        uint      `json:"download_count"`
}

type FileNode struct {
	File
	node.Core
	NodeId uint `json:"node_id"`
}

var (
	SchemaFile schema.Schema
)

func init() {
	awesome_error.CheckFatal(util.InitSchema(&SchemaFile, &File{}))
}
