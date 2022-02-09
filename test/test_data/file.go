package test_data

import (
	"github.com/opencontainers/go-digest"
	"github.com/ssst0n3/dkdk/model"
	"gorm.io/gorm"
	"time"
)

var (
	FileA = model.File{
		Model: gorm.Model{
			ID: 1,
		},
		FileItem: model.FileItem{
			FileIdentifier: model.FileIdentifier{
				Filename: "FileA",
				Digest:   digest.FromString("A"),
			},
			Size: 1000,
		},
		LastModified: time.Now(),
	}
	FileB = model.File{
		Model: gorm.Model{
			ID: 2,
		},
		FileItem: model.FileItem{
			FileIdentifier: model.FileIdentifier{
				Filename: "FileB",
				Digest:   digest.FromString("B"),
			},
			Size: 1000,
		},
		LastModified: time.Now(),
	}
	Files = []model.File{
		FileA,
		FileB,
	}
)
