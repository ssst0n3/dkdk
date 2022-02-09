package test_common

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/awesome_reflect"
	"gorm.io/gorm"
)

func DeleteTable(db *gorm.DB, modelPtr interface{}) (err error) {
	err = db.Migrator().DropTable(modelPtr)
	awesome_error.CheckErr(err)
	return
}

func MakeTableEmpty(db *gorm.DB, modelPtr interface{}) (err error) {
	awesome_reflect.MustPointer(modelPtr)
	err = DeleteTable(db, modelPtr)
	if err != nil {
		return
	}
	err = db.AutoMigrate(modelPtr)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func InitTable(db *gorm.DB, modelPtr interface{}, records interface{}) (err error) {
	awesome_reflect.MustPointer(modelPtr)
	awesome_reflect.MustPointer(records)
	err = MakeTableEmpty(db, modelPtr)
	if err != nil {
		return
	}
	err = db.Create(records).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
