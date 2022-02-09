package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/dkdk/database"
	"github.com/ssst0n3/dkdk/test/test_config"
)

var router = gin.Default()

func init() {
	test_config.Init()
	database.Init()
}
