package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ericpai/please/module/schedule"
	"github.com/ericpai/please/module/task"
	"github.com/ericpai/please/rest/adapter"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskModel     task.Model
	scheduleModel schedule.Model
}

func NewTaskController(taskModel task.Model, scheduleModel schedule.Model) *TaskController {
	return &TaskController{taskModel: taskModel, scheduleModel: scheduleModel}
}

func (t *TaskController) GetAll(c *gin.Context, req *adapter.Request, resp *adapter.Response) {
	pos, err := t.taskModel.GetAll(c.Request.Context())
	if err != nil {
		resp.Meta.Code = http.StatusInternalServerError
		resp.Meta.Message = fmt.Sprintf("内部错误: %s", err.Error())
		return
	}
	for _, v := range pos {
		resp.Data.Tasks = append(resp.Data.Tasks, adapter.TaskModelToExternal(v))
	}
	resp.Meta.Code = http.StatusOK
}

func (t *TaskController) GetByID(c *gin.Context, req *adapter.Request, resp *adapter.Response) {
	idStr := c.Param("task_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Meta.Code = http.StatusBadRequest
		resp.Meta.Message = fmt.Sprintf("ID 格式错误: %s", err.Error())
		return
	}
	po, err := t.taskModel.GetByID(c.Request.Context(), id)
	if err != nil {
		resp.Meta.Code = http.StatusInternalServerError
		resp.Meta.Message = fmt.Sprintf("内部错误: %s", err.Error())
		return
	}
	resp.Data.Tasks = append(resp.Data.Tasks, adapter.TaskModelToExternal(po))
	resp.Meta.Code = http.StatusOK
}

func (t *TaskController) Create(c *gin.Context, req *adapter.Request, resp *adapter.Response) {
	newPO, err := t.taskModel.Insert(c.Request.Context(), adapter.TaskExternalToModel(req.Task))
	if errors.Is(err, task.ErrInvalidParam) {
		resp.Meta.Code = http.StatusUnprocessableEntity
		resp.Meta.Message = err.Error()
		return
	} else if err != nil {
		resp.Meta.Code = http.StatusInternalServerError
		resp.Meta.Message = fmt.Sprintf("内部错误: %s", err.Error())
		return
	}
	if err = t.scheduleModel.Update(int(newPO.ID), newPO.Schedule); err != nil {
		resp.Meta.Code = http.StatusInternalServerError
		resp.Meta.Message = fmt.Sprintf("内部错误: %s", err.Error())
		return
	}
	resp.Data.Tasks = append(resp.Data.Tasks, adapter.TaskModelToExternal(newPO))
	resp.Meta.Code = http.StatusCreated
}

func (t *TaskController) Update(c *gin.Context, req *adapter.Request, resp *adapter.Response) {
	idStr := c.Param("task_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Meta.Code = http.StatusBadRequest
		resp.Meta.Message = fmt.Sprintf("ID 格式错误: %s", err.Error())
		return
	}
	oldPO, err := t.taskModel.GetByID(c.Request.Context(), id)
	if errors.Is(err, task.ErrNotFound) {
		resp.Meta.Code = http.StatusNotFound
		resp.Meta.Message = err.Error()
		return
	}
	if err != nil {
		resp.Meta.Code = http.StatusInternalServerError
		resp.Meta.Message = fmt.Sprintf("内部错误: %s", err.Error())
		return
	}
	reqPO := adapter.TaskExternalToModel(req.Task)
	oldPO.Address = reqPO.Address
	oldPO.Backend = reqPO.Backend
	oldPO.Password = reqPO.Password
	oldPO.SourcePath = reqPO.SourcePath
	oldPO.DestPath = reqPO.DestPath
	oldPO.Enabled = reqPO.Enabled
	oldPO.User = reqPO.User
	oldPO.Schedule = reqPO.Schedule
	newPO, err := t.taskModel.Update(c.Request.Context(), id, oldPO)
	if errors.Is(err, task.ErrInvalidParam) {
		resp.Meta.Code = http.StatusUnprocessableEntity
		resp.Meta.Message = err.Error()
		return
	} else if err != nil {
		resp.Meta.Code = http.StatusInternalServerError
		resp.Meta.Message = fmt.Sprintf("内部错误: %s", err.Error())
		return
	}
	if err = t.scheduleModel.Update(int(newPO.ID), newPO.Schedule); err != nil {
		resp.Meta.Code = http.StatusInternalServerError
		resp.Meta.Message = fmt.Sprintf("内部错误: %s", err.Error())
		return
	}
	resp.Data.Tasks = append(resp.Data.Tasks, adapter.TaskModelToExternal(newPO))
	resp.Meta.Code = http.StatusCreated
}

func (t *TaskController) Delete(c *gin.Context, req *adapter.Request, resp *adapter.Response) {
	idStr := c.Param("task_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Meta.Code = http.StatusBadRequest
		resp.Meta.Message = fmt.Sprintf("ID 格式错误: %s", err.Error())
		return
	}

	if err = t.taskModel.Delete(c.Request.Context(), id); err != nil {
		resp.Meta.Code = http.StatusInternalServerError
		resp.Meta.Message = fmt.Sprintf("内部错误: %s", err.Error())
		return
	}
	if err = t.scheduleModel.Update(int(id), ""); err != nil {
		resp.Meta.Code = http.StatusInternalServerError
		resp.Meta.Message = fmt.Sprintf("内部错误: %s", err.Error())
		return
	}
	resp.Meta.Code = http.StatusNoContent
}
