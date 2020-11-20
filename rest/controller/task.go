package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ericpai/please/module/task"
	"github.com/ericpai/please/rest/adapter"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskModel task.Model
}

func NewTaskController(taskModel task.Model) *TaskController {
	return &TaskController{taskModel: taskModel}
}

func (t *TaskController) GetAll(c *gin.Context, envelope *adapter.Response) {
	pos, err := t.taskModel.GetAll(c.Request.Context())
	if err != nil {
		envelope.Meta.Code = http.StatusInternalServerError
		envelope.Meta.Message = fmt.Sprintf("内部错误: %s", err.Error())
		return
	}
	for _, v := range pos {
		envelope.Data.Tasks = append(envelope.Data.Tasks, adapter.TaskModelToExternal(v))
	}
	envelope.Meta.Code = http.StatusOK
}

func (t *TaskController) GetByID(c *gin.Context, envelope *adapter.Response) {
	idStr := c.Param("task_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		envelope.Meta.Code = http.StatusBadRequest
		envelope.Meta.Message = fmt.Sprintf("ID 格式错误: %s", err.Error())
		return
	}
	po, err := t.taskModel.GetByID(c.Request.Context(), id)
	if err != nil {
		envelope.Meta.Code = http.StatusInternalServerError
		envelope.Meta.Message = fmt.Sprintf("内部错误: %s", err.Error())
		return
	}
	envelope.Data.Tasks = append(envelope.Data.Tasks, adapter.TaskModelToExternal(po))
	envelope.Meta.Code = http.StatusOK
}

func (t *TaskController) Create(c *gin.Context, envelope *adapter.Response) {

}

func (t *TaskController) Update(c *gin.Context, envelope *adapter.Response) {

}

func (t *TaskController) Delete(c *gin.Context, envelope *adapter.Response) {
	idStr := c.Param("task_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		envelope.Meta.Code = http.StatusBadRequest
		envelope.Meta.Message = fmt.Sprintf("ID 格式错误: %s", err.Error())
		return
	}

	if err = t.taskModel.Delete(c.Request.Context(), id); err != nil {
		envelope.Meta.Code = http.StatusInternalServerError
		envelope.Meta.Message = fmt.Sprintf("内部错误: %s", err.Error())
		return
	}
	envelope.Meta.Code = http.StatusNoContent
}
