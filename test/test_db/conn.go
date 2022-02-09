package test_db

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/dkdk/test/test_config"
	"github.com/ssst0n3/lightweight_api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	test_config.Init()
	var err error
	dsn := lightweight_api.GetDsnFromEnvNormal()
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	awesome_error.CheckFatal(err)
}
