package model

import (
	"context"

	"github.com/ericpai/please/module/task"
)

type defaultModel struct {
	taskDB task.DB
}

func New(taskDB task.DB) task.Model {
	return &defaultModel{
		taskDB: taskDB,
	}
}

func (d *defaultModel) GetAll(ctx context.Context) ([]task.PO, error) {
	return d.taskDB.SelectAll(ctx)
}

func (d *defaultModel) GetByID(ctx context.Context, id int) (task.PO, error) {
	return d.taskDB.SelectByID(ctx, uint(id))
}

func (d *defaultModel) Delete(ctx context.Context, id int) error {
	return d.taskDB.Delete(ctx, uint(id))
}
