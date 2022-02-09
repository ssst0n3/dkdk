package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	v1 "github.com/ssst0n3/dkdk/api/v1"
	"github.com/ssst0n3/dkdk/config"
	"github.com/ssst0n3/dkdk/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

type ResponsePing struct {
	Message string `json:"message" example:"pong"`
}

// Ping godoc
// @Summary ping pong
// @Description return pong
// @ID ping-pong
// @Accept  json
// @Produce  json
// @Success 200 {object} ResponsePing
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, ResponsePing{Message: "pong"})
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	// cors
	if config.AllowOrigins != nil {
		corsConfig := cors.DefaultConfig()
		//corsConfig.AllowAllOrigins = true
		corsConfig.AllowCredentials = true
		corsConfig.AllowOrigins = append(corsConfig.AllowOrigins, config.AllowOrigins...)
		router.Use(cors.New(corsConfig))
	}
	// frontend
	{
		router.Use(static.Serve("/", static.LocalFile("./dist", false)))
	}
	// ping pong
	{
		router.GET("/ping", Ping)
	}
	// swagger
	{
		// programmatically set swagger info
		docs.SwaggerInfo.Title = "DKDK API"
		docs.SwaggerInfo.Description = "Netdisk based on docker registry api v2"
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = fmt.Sprintf("127.0.0.1:%s", config.LocalListenPort)
		docs.SwaggerInfo.BasePath = "/"
		docs.SwaggerInfo.Schemes = []string{"http"}
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	{
		v1.InitRouter(router)
	}
	return router
}
