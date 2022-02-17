package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/cipher"
	"github.com/ssst0n3/dkdk/api/v1/repository"
	"github.com/ssst0n3/lightweight_api/example/resource/auth"
	"github.com/ssst0n3/lightweight_api/example/resource/initialize"
	"github.com/ssst0n3/lightweight_api/example/resource/kv_config"
	"github.com/ssst0n3/lightweight_api/example/resource/user"
	"github.com/ssst0n3/lightweight_api/middleware"
	"time"
)

func InitRouter(router *gin.Engine) {
	cipher.Init()
	auth.DurationToken = time.Hour * 24
	auth.InitRouter(router)
	user.InitRouter(router)
	kv_config.InitRouter(router)
	initialize.FlagUseKVConfig = true
	initialize.InitRouter(router)

	repositoryGroup := router.Group(repository.Resource.BaseRelativePath)
	repositoryGroupUser := router.Group(repository.Resource.BaseRelativePath, middleware.JwtUser())
	{
		repositoryGroup.POST("/list", repository.ListFileFromRepositoryConfig)
		repositoryGroup.POST("/download/:digest", repository.DownloadFileFromRepositoryConfig)
		repositoryGroup.POST("/upload", repository.UploadFileFromRepositoryConfig)

		repositoryGroupUser.GET("", repository.ListRepository)
		repositoryGroupUser.POST("", repository.AddRepository)
		repositoryGroupUser.DELETE("/:id", repository.Resource.DeleteResource)
		repositoryGroupUser.GET("/show/:id", repository.ShowRepository)
		repositoryGroupUser.GET("/list", repository.ListFileFromAllRepository)
		repositoryGroupUser.GET("/list/:id", repository.ListFileByRepositoryId)
		repositoryGroupUser.GET("/download/:id/:digest", repository.DownloadFileByRepositoryId)
		repositoryGroupUser.PUT("/:id", repository.UpdateRepository)
		repositoryGroupUser.POST("/upload/:id", repository.UploadFileByRepositoryId)
	}

	taskGroupUser := router.Group(TaskResource.BaseRelativePath, middleware.JwtUser())
	{
		taskGroupUser.GET("", ListTask)
		taskGroupUser.POST("", CreateTask)
		taskGroupUser.POST("/batch/create", BatchCreateTask)
		taskGroupUser.POST("/batch/start", BatchStartTask)
		taskGroupUser.POST("/action/:id", TaskAction)
	}

	fileGroup := router.Group(FileResource.BaseRelativePath, middleware.JwtUser())
	{
		fileGroup.GET("", FileResource.ListResource)
		fileGroup.GET("/dir/:id", ListFileUnderDir)
		fileGroup.GET("/exists", CheckFilenameAlreadyExists)
		fileGroup.POST("/repository/:id", UploadRepositoryFileToDirectory)
	}

	directoryGroup := router.Group(DirectoryResource.BaseRelativePath, middleware.JwtUser())
	{
		directoryGroup.GET("", DirectoryResource.ListResource)
		directoryGroup.GET("/list/:id", ListDirectoryUnderDir)
		directoryGroup.GET("/path/:id", GetDirectoryPath)
		directoryGroup.POST("/node", CreateDirectoryNode)
		directoryGroup.DELETE("/:id", DirectoryResource.DeleteResource)
	}

	nodeGroup := router.Group(NodeResource.BaseRelativePath, middleware.JwtUser())
	{
		nodeGroup.DELETE("/:id", NodeResource.DeleteResource)
		nodeGroup.DELETE("", BatchDeleteNode)
		nodeGroup.PUT("/:id", Move)
	}
}
