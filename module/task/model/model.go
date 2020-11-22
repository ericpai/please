package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ericpai/please/module/task"
	"gorm.io/gorm"
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
	res, err := d.taskDB.SelectByID(ctx, uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, fmt.Errorf("%w: %d", task.ErrNotFound, id)
	}
	return res, err
}

func (d *defaultModel) Delete(ctx context.Context, id int) error {
	return d.taskDB.Delete(ctx, uint(id))
}

func (d *defaultModel) Insert(ctx context.Context, v task.PO) (task.PO, error) {
	v.ID = 0
	c := time.Now()
	v.CreatedTime = c
	v.UpdatedTime = c
	if err := d.validate(v); err != nil {
		return v, err
	}
	return d.taskDB.Insert(ctx, v)
}
func (d *defaultModel) Update(ctx context.Context, id int, v task.PO) (task.PO, error) {
	if err := d.validate(v); err != nil {
		return v, err
	}
	return d.taskDB.Update(ctx, v)
}

func (d *defaultModel) validate(v task.PO) error {
	if v.Address == "" {
		return fmt.Errorf("%w: address", task.ErrInvalidParam)
	}
	if v.Backend != task.BackendLinux && v.Backend != task.BackendWindows {
		return fmt.Errorf("%w: backend", task.ErrInvalidParam)
	}
	if v.SourcePath == "" {
		return fmt.Errorf("%w: sourcePath", task.ErrInvalidParam)
	}
	if v.DestPath == "" {
		return fmt.Errorf("%w: destPath", task.ErrInvalidParam)
	}
	return nil
}
