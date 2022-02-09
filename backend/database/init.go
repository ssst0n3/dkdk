package database

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/lightweight_api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(lightweight_api.GetDsnFromEnvNormal()), &gorm.Config{})
	lightweight_api.DB = DB
	awesome_error.CheckFatal(Migrate())
	awesome_error.CheckFatal(err)
}
