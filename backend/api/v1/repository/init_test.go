package repository

import (
	"dkdk/database"
	"dkdk/test/test_config"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func init() {
	test_config.Init()
	database.Init()
}
