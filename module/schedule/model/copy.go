package model

import (
	"context"
	"log"
	"time"

	"github.com/ericpai/please/module/task"
	"github.com/ericpai/please/module/worker"
)

func RemoteCopy(id int, taskModel task.Model, workerModel worker.Model) {
	ctx := context.Background()
	po, err := taskModel.GetByID(ctx, id)
	if err != nil {
		log.Printf("执行定时任务[RemoteCopy]错误：%s\n", err.Error())
		return
	}
	if !po.Enabled {
		return
	}
	workCtx := worker.Context{
		Address:    po.Address,
		User:       po.User,
		Password:   po.Password,
		DestPath:   po.DestPath,
		SourcePath: po.SourcePath,
	}
	succeed := true
	if err = workerModel.Copyfile(ctx, workCtx); err != nil {
		log.Printf("执行定时任务[RemoteCopy]: %s\n", err.Error())
		succeed = false
	}
	po.Succeed = succeed
	po.UpdatedTime = time.Now()
	if _, err = taskModel.Update(ctx, int(po.ID), po); err != nil {
		log.Printf("更新任务状态[RemoteCopy]: %s\n", err.Error())
	}
}
