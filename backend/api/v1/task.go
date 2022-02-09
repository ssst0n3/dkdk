package v1

import (
	"dkdk/database"
	"dkdk/model"
	"dkdk/task/offline_download"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/response"
	"net/http"
	"strconv"
)

const (
	TaskActionStart  = "start"
	TaskResourceName = "task"
)

var TaskResource = lightweight_api.NewResource(
	TaskResourceName, model.SchemaTask.Table, model.Task{}, "",
)

func ListTask(c *gin.Context) {
	userId, err := lightweight_api.GetUserId(c)
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 0
	}
	sizeStr := c.Query("size")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 0
	}
	tasks, err := database.ListTaskByUserId(userId, uint(page), uint(size))
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var task *model.Task
	TaskResource.CreateResourceTemplate(c, func(modelPtr interface{}) (err error) {
		task = modelPtr.(*model.Task)
		userId, err := lightweight_api.GetUserId(c)
		if err != nil {
			lightweight_api.HandleStatusBadRequestError(c, err)
			return
		}
		task.UserId = userId
		return
	}, func(id uint) (err error) {
		// todo: routine
		//spew.Dump(task)
		//task.ID = id
		offline_download.ChanTask <- *task
		return
	})
}

func BatchCreateTask(c *gin.Context) {
	// TODO: set max size of json
	var tasks []model.Task
	err := c.BindJSON(&tasks)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	userId, err := lightweight_api.GetUserId(c)
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	err = database.BatchCreateTask(tasks, userId)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	c.Status(http.StatusOK)
}

func TaskAction(c *gin.Context) {
	_, resource, err := TaskResource.MustResourceExistsGetModelByIdAutoParseParam(c)
	if err != nil {
		return
	}
	task := resource.(model.Task)
	action := c.Query("action")
	switch action {
	case TaskActionStart:
		offline_download.ChanTask <- task
		response.Success200(c, fmt.Sprintf("task %s start", task.OriginUrl))
		return
	}
}

func BatchStartTask(c *gin.Context) {

}
