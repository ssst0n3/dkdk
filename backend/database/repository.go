package database

import (
	"dkdk/model"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

func ListRepositoryForResponse(userId uint) (result []model.RepositoryConfigResponse, err error) {
	//whereQuery := awesome_libs.Format(
	//	"{.user_id}=?", awesome_libs.Dict{
	//		"user_id": model.SchemaRepository.FieldsByName["UserId"].DBName,
	//	})
	err = DB.Model(&model.Repository{UserId: userId}).Find(&result).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func GetRepositoryById(id uint) (repository model.Repository, err error) {
	err = DB.Model(&model.Repository{}).First(&repository, id).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func ListRepositoryByUserId(id uint) (repository []model.Repository, err error) {
	err = DB.Where(model.Repository{
		UserId: id,
	}).Find(&repository).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func CheckUserHasRepository(userId uint, repositoryId uint) (has bool, err error) {
	var repository model.Repository
	err = DB.Model(&repository).First(&repository, repositoryId).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	has = repository.UserId == userId
	return
}
