package test_data

import (
	"dkdk/model"
	"gorm.io/gorm"
)

var (
	DirectoryA = model.Directory{
		Model: gorm.Model{
			ID: 1,
		},
		Filename: "DirA",
	}
	DirectoryB = model.Directory{
		Model: gorm.Model{
			ID: 2,
		},
		Filename: "DirB",
	}
	Directories = []model.Directory{
		DirectoryA,
		DirectoryB,
	}
)
