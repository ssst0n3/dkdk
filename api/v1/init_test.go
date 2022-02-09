package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/dkdk/database"
	"github.com/ssst0n3/dkdk/test/test_config"
)

var router = gin.Default()

func init() {
	test_config.Init()
	InitRouter(router)
	//lightweight_api.InitConnector()
	database.Init()
}
