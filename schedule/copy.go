package schedule

import (
	"context"
	"log"

	"github.com/ericpai/please/module/task"
	"github.com/ericpai/please/module/worker"
)

func RemoteCopy(taskModel task.Model, workerModel worker.Model) {
	ctx := context.Background()
	pos, err := taskModel.GetAll(ctx)
	if err != nil {
		log.Printf("执行定时任务[RemoteCopy]错误：%s\n", err.Error())
		return
	}
	for _, po := range pos {
		if !po.Enabled {
			continue
		}
		workCtx := worker.Context{
			Address:    po.Address,
			User:       po.User,
			Password:   po.Password,
			DestPath:   po.DestPath,
			SourcePath: po.SourcePath,
		}
		if err := workerModel.Copyfile(ctx, workCtx); err != nil {
			log.Printf("执行定时任务[RemoteCopy]: %s\n", err.Error())
		}
	}
}
